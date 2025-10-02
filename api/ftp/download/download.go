package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

type info struct {
	UF  string `json:"uf"`
	Ano string `json:"ano"`
	Mes string `json:"mes"`
}

func main() {
	var infos = []info{
		{UF: "SC", Ano: "25", Mes: "05"},
		{UF: "SC", Ano: "25", Mes: "04"},
		{UF: "SC", Ano: "25", Mes: "03"},
	}

	for i := range infos {
		fmt.Println(infos[i])
	}

	path := "ftp.datasus.gov.br:21"
	login, password := "anonymous", "anonymous"

	resp, err := ConnectFtp(path, login, password)
	if err != nil {
		panic(err)
	}

	ftp_path := "/dissemin/publicos/CNES/200508_/Dados/ST/STSC2506.dbc"
	download_path := "api/ftp/archives/"
	name := "teste.dbc"
	err = ReadAndDownload(resp, download_path, name, ftp_path)
	if err != nil {
		panic(err)
	}

}

func ConnectFtp(path string, login string, password string) (response *ftp.ServerConn, err error) {
	resp, err := ftp.Dial(path, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = resp.Login(login, password)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

func ReadAndDownload(resp *ftp.ServerConn, download_path string, name string, ftp_path string) error {
	read, err := resp.Retr(ftp_path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(download_path)
	fmt.Println(err)
	if err != nil {
		err = os.MkdirAll(download_path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	archive, err := os.Create(download_path + name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(archive, read)

	defer archive.Close()
	defer read.Close()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
