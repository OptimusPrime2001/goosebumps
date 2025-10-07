package main

import "fmt"

func TestOperator() {
	a, b := 4, 15
	tong := a + b
	hieu := b - a
	tich := a * b
	thuong := float32(b) / float32(a)
	fmt.Printf("Tong cua %d va %d la: %d\n", a, b, tong)
	fmt.Printf("Hieu cua %d va %d la: %d\n", b, a, hieu)
	fmt.Printf("Tich cua %d va %d la: %d\n", a, b, tich)
	fmt.Printf("Thuong cua %d va %d la: %g \n", b, a, thuong)
}
