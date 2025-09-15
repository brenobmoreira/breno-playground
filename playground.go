package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/api"
)

func main() {

	var cep string

	fmt.Println("Digite seu CEP: ")
	fmt.Scanln(&cep)
	cepBody, err := api.Cep(cep)

	if err != nil {
		panic(err)
	}

	fmt.Println(cepBody)
}
