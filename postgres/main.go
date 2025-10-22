package main

import (
	"fmt"

	"github.com/brenobmoreira/breno-playground/postgres/functions"
)

func main() {

	var categoria = []functions.Categoria{
		{Nome: "Breno"},
		{Nome: "Felipe"},
		{Nome: "Jose"},
	}

	var lista_ids = []int64{}

	for name := range categoria {
		id, err := functions.Insert(categoria[name])
		if err != nil {
			panic(err)
		}
		lista_ids = append(lista_ids, id)
	}

	cat, err := functions.Get(lista_ids[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("Select id[0]: ", cat)

	categoria_update := functions.Categoria{Nome: "Joao"}
	updated, err := functions.Update(lista_ids[0], categoria_update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Row updated: ", updated)

	deleted, err := functions.Delete(lista_ids[1])
	if err != nil {
		panic(err)
	}
	fmt.Println("Row deleted: ", deleted)

	sc, err := functions.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("Select all ids: ", sc)
}
