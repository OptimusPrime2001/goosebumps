package main

import (
	"fmt"
	"time"
)

func TestSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// goroutine gửi dữ liệu sau 1s
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	// goroutine gửi dữ liệu sau 2s
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	// Nếu không dùng Select, phải chờ cả 5s để nhận được dữ liệu từ ch1, rồi mới nhận được ch2
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)

	// Dùng Select, có thể nhận được dữ liệu từ ch2 trước ch1, vì ch2 có delay 2s
	for range 2 {
		time.Sleep(2 * time.Second)
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:
			fmt.Println("No message yet")
		}

	}
}
