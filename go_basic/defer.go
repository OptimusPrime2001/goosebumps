package main

import "fmt"

// Các defer được gọi theo thứ tự LIFO (Last In First Out)
func TestDefer() {
	//defer nào được khai báo sau sẽ được thực thi trước — giống như một ngăn xếp (stack).
	fmt.Println("Test Defer start running")

	defer fmt.Println("Dòng này bị block cho đến khi hàm chạy xong")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Test Defer running end")
}
