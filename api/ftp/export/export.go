package main

import (
	"fmt"
	"os/exec"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
	"github.com/brenobmoreira/breno-playground/api/csv/writer"
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

func ReadChan(test chan []string) {
	fmt.Println("Testing how works")
	fmt.Println(<-test)
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

	var line uint32
	var estabString [][]string

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

		data := []string{
			p.ID,
			p.CodigoMunicipio,
			p.CodigoCEP,
			p.CPFouCNPJ,
		}

		test := make(chan []string)
		go ReadChan(test)
		test <- data

		estabString = append(estabString, data)

		// fmt.Printf("EstabelecimentoDBF: %+v \n", p)

	}

	output := "/home/dev/playground/breno-playground/api/ftp/write.csv"
	err = writer.WriteCsv(estabString, output)
	if err != nil {
		panic(err)
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
