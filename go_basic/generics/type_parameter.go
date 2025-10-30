package main

import (
	"cmp"
	"fmt"
)

func CountValue[T comparable](s []T, x T) int {
	result := 0
	for _, v := range s {
		if v == x {
			result++
		}
	}
	return result

}
func compare[T cmp.Ordered](a, b T) bool {
	return a < b
}
func TestGenerics() {
	fmt.Println(compare(1, 2))
	test1 := []int{1, 2, 56, 324, 4, 56, 54657}
	fmt.Println(CountValue(test1, 56))
	test2 := []string{"1", "2", "56", "324", "4", "56", "54657"}
	fmt.Println(CountValue(test2, "56"))
}
