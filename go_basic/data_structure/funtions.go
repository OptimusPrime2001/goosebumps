package main

import (
	"fmt"
	"math"
)

func closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(4, 5)
}
func TestFunc() {
	hypol := func(x, y float64) float64 {
		return x * y
	}
	test1, test2 := closure(), closure()
	var (
		value1 = [2]int{test1(4), test1(6)}
		value2 = [2]int{test2(3), test2(5)}
	)

	fmt.Println("===========Test Functions")
	fmt.Println(value1, value2)
	fmt.Println(hypol(5, 6))
	fmt.Println(compute(hypol))
	fmt.Println(compute(math.Pow))
}
