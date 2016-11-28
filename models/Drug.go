package models

import "github.com/jinzhu/gorm"

//Drug drug information
type Drug struct {
	gorm.Model
	Name                  string `json:"name"`
	Indication            string `json:"indication"`
	Classification        string `json:"classification"`
	DosageFormAndStrength string `json:"dosage_form_and_strength"`
	AdverseDrugEffects    string `json:"adverse_drug_effects"`
	DrugInteractions      string `json:"drug_interactions"`
	Precaution            string `json:"precaution"`
	Contraindications     string `json:"contraindications"`
	PregnancyCategory     string `json:"pregnancy_category"`
	DosageRegimen         string `json:"dosage_regimen"`
	FoodIntake            string `json:"food_intake"`
}

//SetName sets drug name
func (d *Drug) SetName(name string) {
	d.Name = name
}

//SetIndication sets drug name
func (d *Drug) SetIndication(indication string) {
	d.Indication = indication
}

//SetClassification sets drug name
func (d *Drug) SetClassification(classification string) {
	d.Classification = classification
}

//SetDosageFormAndStrength sets drug name
func (d *Drug) SetDosageFormAndStrength(dosageFormStrength string) {
	d.DosageFormAndStrength = dosageFormStrength
}

//SetAdverseDrugEffects sets drug name
func (d *Drug) SetAdverseDrugEffects(adverseDrugEffects string) {
	d.AdverseDrugEffects = adverseDrugEffects
}

//SetDrugInteractions sets drug name
func (d *Drug) SetDrugInteractions(drugInteractions string) {
	d.DrugInteractions = drugInteractions
}

//SetPrecaution sets drug name
func (d *Drug) SetPrecaution(precaution string) {
	d.Precaution = precaution
}

//SetContraindications sets drug name
func (d *Drug) SetContraindications(contraindications string) {
	d.Contraindications = contraindications
}

//SetPregnancyCategory sets drug name
func (d *Drug) SetPregnancyCategory(pregnancyCategory string) {
	d.PregnancyCategory = pregnancyCategory
}

//SetDosageRegimen sets drug name
func (d *Drug) SetDosageRegimen(dosageRegimen string) {
	d.DosageRegimen = dosageRegimen
}

//SetFoodIntake sets drug name
func (d *Drug) SetFoodIntake(foodIntake string) {
	d.FoodIntake = foodIntake
}
