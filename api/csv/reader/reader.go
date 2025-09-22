package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("output/EQSC2506.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	for i := 0; i < 5; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
