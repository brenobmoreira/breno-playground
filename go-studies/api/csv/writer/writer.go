package writer

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func WriteCsv(records [][]string, path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return err
}
