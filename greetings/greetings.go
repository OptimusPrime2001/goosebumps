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

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Greetings(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
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
