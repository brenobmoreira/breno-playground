package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/api/csv/reader"
)

func main() {
	path := "output/EQSC2506.csv"

	err := reader.ReadCsv(path)
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
	}
}
