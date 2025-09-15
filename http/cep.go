package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	var cep string

	fmt.Println("Enter your CEP: ")
	fmt.Scanln(&cep)

	url := "https://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	bodyBites, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBites)

	fmt.Println(bodyString)

}
