package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/api/csv/reader"
	"github.com/brenobmoreira/breno-playground/api/csv/writer"
)

func main() {
	path := []string{"EQSC2506.csv", "STSC2506.csv"}

	for j := range path {
		records, err := reader.ReadCsv("output/" + path[j])
		if err != nil {
			fmt.Printf("Error reading CSV: %v\n", err)
		}

		for i := range records {
			fmt.Println("--  Linha", i, " --")
			fmt.Println(records[i])
		}

		output := "api/csv/output/" + path[j]
		err = writer.WriteCsv(records, output)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Sucessufully wrote %s\n", output)
	}
}
