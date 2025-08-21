package main

import "fmt"

func CountValue[T comparable](s []T, x T) int {
	result := 0
	for _, v := range s {
		if v == x {
			result++
		}
	}
	return result

}
func TestGenerics() {
	test1 := []int{1, 2, 56, 324, 4, 56, 54657}
	fmt.Println(CountValue(test1, 56))
	test2 := []string{"1", "2", "56", "324", "4", "56", "54657"}
	fmt.Println(CountValue(test2, "56"))
}
