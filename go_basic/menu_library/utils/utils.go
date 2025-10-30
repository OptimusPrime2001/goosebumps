package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func GetPostiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		value, error := strconv.Atoi(input)
		if error == nil && value > 0 {
			return value
		}
		fmt.Println("Invalid input. Please enter a positive integer.")
	}
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
