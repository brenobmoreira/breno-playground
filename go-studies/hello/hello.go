package main

import (
	"fmt"
	"log"

	"github.com/brenobmoreira/breno-playground/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Didico", "Burigo", "Coisa"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
