package main

import "fmt"

func genElement() func() int {
	init, next := 0, 1

	return func() int {
		value := init
		init, next = next, next+init
		// init = next
		// next = init + next
		return value
	}
}
func Finbonaci() {
	element := genElement()
	fmt.Println("============Finbonaci==========")

	for range 10 {
		fmt.Println(element())
	}
}
