package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/api/csv/reader"
	"github.com/brenobmoreira/breno-playground/api/csv/writer"
)

func main() {
	path := "output/EQSC2506.csv"

	records, err := reader.ReadCsv(path)
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
	}

	for i := range records {
		fmt.Println("--  Linha", i, " --")
		fmt.Println(records[i])
	}

	output := "api/csv/output/file.csv"
	err = writer.WriteCsv(records, output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sucessufully wrote %s\n", output)

}
