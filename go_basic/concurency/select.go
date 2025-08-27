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
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	// goroutine gửi dữ liệu sau 2s
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	for range 2 {
		time.Sleep(1 * time.Second)
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
