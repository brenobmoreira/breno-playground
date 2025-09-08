package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Didico")
	fmt.Println(message)
}
