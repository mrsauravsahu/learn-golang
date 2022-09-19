package main

import (
  "fmt"

  "mrsauravsahu.in/greetings"
)

func main() {
  // Get a message and print it
  message := greetings.Hello("Sahu")
  fmt.Println(message)
}

