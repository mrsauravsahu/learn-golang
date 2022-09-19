package main

import (
	"fmt"
	"log"

	"mrsauravsahu.in/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a message and print it
	message, err := greetings.Hello("Saurav")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
