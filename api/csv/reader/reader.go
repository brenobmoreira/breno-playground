package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadCsv(path string) error {
	file, err := os.Open(path)
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

	return err
}
