package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/postgres/functions"
)

func main() {

	var categoria = functions.Categoria{ID: 1, Nome: "Breno"}

	id, err := functions.Insert(categoria)
	if err != nil {
		panic(err)
	}

	cat, err := functions.Get(id)
	if err != nil {
		panic(err)
	}
	fmt.Print(cat)
}
