package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Greetings(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "Has error", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi there, i'm %v, nice to meet you.", name)
	return message, nil
}
func GetAge(age int) int {
	result := age + 2001
	return result
}
