package main

import "fmt"

func TestPointer() {
	x, y := 20, 2025
	var p, q *int

	// khai báo pointer p là cho địa chỉ ô nhớ lưu lưu trữ giá trị của x
	p = &x

	//Thay đổi giá trị địa chỉ ô nhớ qua con trỏ p
	*p = 2003
	fmt.Println(p, x, y)

	q = &y
	*q /= 25
	fmt.Println(y)
}
