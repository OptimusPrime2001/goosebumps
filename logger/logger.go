package main

import (
	"encoding/json"
	"log"
	"trung_go/greetings"

	"fmt"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	result, _ := json.MarshalIndent(message, "", "  ")
	fmt.Println("Run module logger: ", string(result), err)
}
