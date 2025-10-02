package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/api/ftp"
)

func main() {

	var cep string

	fmt.Println("Digite seu CEP: ")
	fmt.Scanln(&cep)
	cepBody, err := ftp.Cep(cep)

	if err != nil {
		panic(err)
	}

	fmt.Println(cepBody)
}
