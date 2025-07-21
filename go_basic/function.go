package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func calculateAge(age int) int {
	return 2025 - age
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(add(2, 3), calculateAge(2001))
	a, b := swap("Trung Le", "Họ và tên :")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
