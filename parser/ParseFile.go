package parser

import (
	"log"
	"os"
)

//ParseFile parses the file
func ParseFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln("File " + file + " can't be opened")
	}
	CSVParser(f)
}
