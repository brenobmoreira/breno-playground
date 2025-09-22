package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("api/csv/writer/archive.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	records := [][]string{
		{"Breno Moreira", "blabla@gmail.com"},
		{"Celso Russomano", "blabla2@gmail.com"},
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}
}
