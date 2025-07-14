package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Greetings(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi there, i'm %v, nice to meet you.", name)
    return message
}
func GetAge(age int) int {
	result := age + 2001
	return result
}