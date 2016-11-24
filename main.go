package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("drug_database.csv")
	check(err)
	log.Println(file)
}
