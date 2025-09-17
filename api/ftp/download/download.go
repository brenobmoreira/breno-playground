package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {

	type info struct {
		UF  string `json:"uf"`
		Ano string `json:"ano"`
		Mes string `json:"mes"`
	}

	var infos = []info{
		{UF: "SC", Ano: "25", Mes: "05"},
		{UF: "SC", Ano: "25", Mes: "04"},
		{UF: "SC", Ano: "25", Mes: "03"},
	}

	for i := range infos {
		fmt.Println(infos[i])
	}

	resp, err := ftp.Dial("ftp.datasus.gov.br:21", ftp.DialWithTimeout(5*time.Second))

	if err != nil {
		log.Fatal(err)
	}

	err = resp.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	read, err := resp.Retr("/dissemin/publicos/CNES/200508_/Dados/ST/STSC2403.dbc")
	if err != nil {
		log.Fatal(err)
	}

	archive, err := os.Create("api/ftp/teste.dbc")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(archive, read)

	defer archive.Close()
	defer read.Close()

	if err != nil {
		log.Fatal(err)
	}

}
