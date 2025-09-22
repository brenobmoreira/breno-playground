package reader

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var records [][]string
	r := csv.NewReader(file)

	for i := 0; i < 5; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		records = append(records, record)

		// fmt.Println(record)
	}

	return records, err
}
