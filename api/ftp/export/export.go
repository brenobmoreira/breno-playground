package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

type EstabelecimentoDBF struct {
	ID              string `dbase:"CNES"`
	CodigoMunicipio string `dbase:"CODUFMUN"`
	CodigoCEP       string `dbase:"COD_CEP"`
	CPFouCNPJ       string `dbase:"CPF_CNPJ"`
}

type Estabelecimento struct {
	Cidade string
	Estado string
}

func (e EstabelecimentoDBF) ToEstabelecimento() Estabelecimento {
	var Estado string
	var Cidade string

	if e.CodigoCEP == "420005" {
		Estado = "Santa Catarina"
		Cidade = "Florianopolis"
	}

	return Estabelecimento{
		Cidade: Cidade,
		Estado: Estado,
	}
}

func main() {
	dbc_path := "/home/dev/playground/breno-playground/api/ftp/export/teste.dbc"
	dbf_path := "/home/dev/playground/breno-playground/api/ftp/export/teste.dbf"
	blast := "./blast-dbf"
	dir := "/home/dev/playground/breno-playground/api/ftp/export/"
	err := DBCtoDBF(dbc_path, dbf_path, blast, dir)
	if err != nil {
		panic(err)
	}

	table, err := ReadDbf(dbf_path)
	if err != nil {
		panic(err)
	}
	defer table.Close()

	path := "/home/dev/playground/breno-playground/api/ftp/write.csv"
	directory := filepath.Dir(path)
	if directory != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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

		estabChan := make(chan EstabelecimentoDBF)
		go WriteChan(file, estabChan)
		estabChan <- p

		fmt.Printf("EstabelecimentoDBF: %+v \n", p)

	}
}

func WriteChan(file *os.File, data chan EstabelecimentoDBF) {
	w := csv.NewWriter(file)
	defer w.Flush()
	r := <-data
	record := []string{
		r.ID,
		r.CodigoMunicipio,
		r.CodigoCEP,
		r.CPFouCNPJ}
	if err := w.Write(record); err != nil {
		panic(err)
	}
	fmt.Println(record)
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

func RowToStruct(row *dbase.Row) (EstabelecimentoDBF, error) {
	p := &EstabelecimentoDBF{}
	err := row.ToStruct(p)
	if err != nil {
		return EstabelecimentoDBF{}, err
	}

	return *p, nil
}

func DBCtoDBF(dbc string, dbf string, blast string, dir string) error {
	cmd := exec.Command(blast, dbc, dbf)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running blast-dbf: %v\nOutput: %s\n", err, string(output))
		return err
	}

	fmt.Println("Successfully converted .dbc to .dbf")
	return err
}
