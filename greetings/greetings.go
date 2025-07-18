package greetings

import (
	"errors"
	"fmt"
	"math/rand"
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
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}
func randomGreeting() string {
	listGreeting := []string{
		"Hello guys, my name is %v",
		"Nice to meet you, i'm %v",
		"Is is my pleasure, you can call me %v",
	}
	index := rand.Intn(len(listGreeting))
	return listGreeting[index]
}
func GetAge(age int) int {
	result := age + 2001
	return result
}
