package main

import (
	"os"

	parser "github.com/chaofanman/pharmacopy-rest/parser"
	utils "github.com/chaofanman/pharmacopy-rest/utils"
)

func main() {
	f, err := os.Open("input_files/db3.csv")
	utils.Check(err)

	parser.CSVParser(f)
}
