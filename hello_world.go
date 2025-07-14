package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"trung_go/greetings"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Hello())
		fmt.Print(greetings.GetAge(24))
}