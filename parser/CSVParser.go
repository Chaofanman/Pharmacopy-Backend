package parser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/chaofanman/pharmacopy-rest/database"
	"github.com/chaofanman/pharmacopy-rest/models"
)

//CSVParser parses given CSV file
func CSVParser(f *os.File) {
	db := database.Init()
	fmt.Println("Database in CSV Parser: ", db)
	r := csv.NewReader(bufio.NewReader(f))
	var Drug models.Drug
	var drugs []models.Drug
	for {
		record, err := r.Read()

		if err == io.EOF {
			for i := range drugs {
				db.Create(&drugs[i])
				db.Exec("DELETE FROM drugs WHERE id IN (SELECT id FROM (SELECT id, ROW_NUMBER() OVER (partition BY name, indication, classification, dosage_form_and_strength, adverse_drug_effects, precaution, contraindications, pregnancy_category, dosage_regimen, food_intake ORDER BY id) AS rnum FROM drugs) t WHERE t.rnum > 1);")
			}
			break
		}
		for value := range record {
			switch value {
			case 0:
				Drug.SetName(record[value])
			case 1:
				Drug.SetIndication(record[value])
			case 2:
				Drug.SetClassification(record[value])
			case 3:
				Drug.SetDosageFormAndStrength(record[value])
			case 4:
				Drug.SetAdverseDrugEffects(record[value])
			case 5:
				Drug.SetDrugInteractions(record[value])
			case 6:
				Drug.SetPrecaution(record[value])
			case 7:
				Drug.SetContraindications(record[value])
			case 8:
				Drug.SetPregnancyCategory(record[value])
			case 9:
				Drug.SetDosageRegimen(record[value])
			case 10:
				Drug.SetFoodIntake(record[value])
				drugs = append(drugs, Drug)
			default:
				continue
			}
		}
	}
}
