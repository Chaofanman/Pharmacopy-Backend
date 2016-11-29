package models

//QueuedDrug Where SQL query is stored
type QueuedDrug struct {
	ID   int
	Name string
}

//Drugs is a slice of QueuedDrugs
type Drugs []QueuedDrug

func (slice Drugs) Len() int {
	return len(slice)
}

func (slice Drugs) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice Drugs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
