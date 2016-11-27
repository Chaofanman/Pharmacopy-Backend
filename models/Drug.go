package models

//Drug drug information
type Drug struct {
	Name                  string `json:"name"`
	Indication            string `json:"indication"`
	Classification        string `json:"classification"`
	DosageFormAndStrength string `json:"dosageformandstrength"`
	AdverseDrugEffects    string `json:"adversedrugeffects"`
	DrugInteractions      string `json:"druginteractions"`
	Precaution            string `json:"precaution"`
	Contraindications     string `json:"contraindications"`
	PregnancyCategory     string `json:"pregnancycategory"`
	DosageRegimen         string `json:"dosageregimen"`
	FoodIntake            string `json:"foodintake"`
}
