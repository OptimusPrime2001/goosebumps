package main

import (
	"fmt"
	"time"
)

func TestSwitch() {
	switch age := 2000; age {
	case 2001:
		fmt.Println(2001)
	case 2000:
		fmt.Println(2000)
	default:
		fmt.Println("Default value")
	}
}
func SayHello() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
