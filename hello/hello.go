package main

import (
	"fmt"
	"log"

	"pewpew.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry prefix
	// and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Wonsang", "James", "David", "Nancy", "Dave"}
	// Request a greeting message
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to console and exit
	if err != nil {
		log.Fatal(err)
	}

	// If no error, print message to console
	fmt.Println(messages)
}
