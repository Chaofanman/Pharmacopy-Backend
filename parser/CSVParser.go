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
	// var drugs []models.Drug
	for {
		record, err := r.Read()

		if err == io.EOF {
			// for drug := range drugs {
			// 	fmt.Println(drugs[drug].Name)
			// }
			break
		}
		for value := range record {
			// var Drug models.Drug
			switch value {
			case 0:
				// Drug.Name = record[value]
				Drug.SetName(record[value])
			case 1:
				// Drug[count].Indication = record[value]
				Drug.SetIndication(record[value])
			case 2:
				// Drug[count].Classification = record[value]
				Drug.SetClassification(record[value])
			case 3:
				// Drug[count].DosageFormAndStrength = record[value]
				Drug.SetDosageFormAndStrength(record[value])
			case 4:
				// Drug[count].AdverseDrugEffects = record[value]
				Drug.SetAdverseDrugEffects(record[value])
			case 5:
				// Drug[count].DrugInteractions = record[value]
				Drug.SetDrugInteractions(record[value])
			case 6:
				// Drug[count].Precaution = record[value]
				Drug.SetPrecaution(record[value])
			case 7:
				// Drug[count].Contraindications = record[value]
				Drug.SetContraindications(record[value])
			case 8:
				// Drug[count].PregnancyCategory = record[value]
				Drug.SetPregnancyCategory(record[value])
			case 9:
				// Drug[count].DosageRegimen = record[value]
				Drug.SetDosageRegimen(record[value])
			case 10:
				Drug.SetFoodIntake(record[value])
				db.Create(&Drug)
				// fmt.Println(err)
			default:

			}
			// fmt.Println(value, ": ", record[value])
		}
	}

}
