package main

import "fmt"

func TestBufferedChannel() {
	ch := make(chan int,2)
	ch <- 1
	fmt.Println("Đã gửi 1 vào channel")
	ch <- 2
	fmt.Println("Đã gửi 2 vào channel")
	fmt.Println(<-ch)
	fmt.Println("Đã nhận 1 từ channel")
	fmt.Println(<-ch)
	fmt.Println("Đã nhận 2 từ channel")
	fmt.Println("Kết thúc")
}
