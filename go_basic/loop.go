package main

import "fmt"

func Loop2() {
	a := 1
	for a < 30 {
		a += a
	}
	fmt.Println(a)
}
func Loop() {
	sum := 2001
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
