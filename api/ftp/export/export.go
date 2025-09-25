package main

import (
	"fmt"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

type Estabelecimento struct {
	ID              string `dbase:"CNES"`
	CodigoMunicipio string `dbase:"CODUFMUN"`
	CodigoCEP       string `dbase:"COD_CEP"`
	CPFouCNPJ       string `dbase:"CPF_CNPJ"`
}

func main() {
	path := "/home/dev/playground/breno-playground/api/ftp/export/output.dbf"
	table, err := ReadDbf(path)
	if err != nil {
		panic(err)
	}
	defer table.Close()

	var line uint32

	for !table.EOF() {
		line++
		row, err := table.Next()
		if err != nil {
			panic(err)
		}

		if row.Deleted {
			fmt.Printf("Deleted row at position: %v \n", row.Position)
			continue
		}

		p, err := RowToStruct(row)
		if err != nil {
			fmt.Printf("Error in row: %d, %v", line, err)
			continue
		}

		fmt.Printf("Estabelecimento: %+v \n", p)
	}
}

func ReadDbf(path string) (*dbase.File, error) {
	table, err := dbase.OpenTable(&dbase.Config{
		Filename:   path,
		TrimSpaces: true,
		Untested:   true,
	})
	if err != nil {
		return nil, err
	}

	return table, nil
}

func RowToStruct(row *dbase.Row) (Estabelecimento, error) {
	p := &Estabelecimento{}
	err := row.ToStruct(p)
	if err != nil {
		return Estabelecimento{}, err
	}

	return *p, nil
}
