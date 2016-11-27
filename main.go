package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("InputFiles/db3.csv")
	check(err)

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
