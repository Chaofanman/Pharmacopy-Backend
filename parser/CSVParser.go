package parser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//CSVParser parses given CSV file
func CSVParser(f *os.File) {
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		// fmt.Println(record)
		// fmt.Println(len(record))
		for value := range record {
			fmt.Printf("%d:  %v\n", value, record[value])
		}
	}
}
