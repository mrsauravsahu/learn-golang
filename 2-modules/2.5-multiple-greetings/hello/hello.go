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
	names := []string{"Saurav", "John", "Jerry", "Sarah", "Lisa"}
	messageMap, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messageMap {
		fmt.Printf("%v\n\n%v\n===", name, message)
		fmt.Println()
	}
}
