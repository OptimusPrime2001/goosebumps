package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // gửi tổng vào channel
		fmt.Println("Đã gửi vào channel")
}

func TestChannel() {
    s := []int{7, 2, 8, -9, 4, 0}
		c:=make(chan int)
    go sum(s[:len(s)/2], c) 
    go sum(s[len(s)/2:], c) 
		fmt.Println("Đã gửi 2 goroutine")
    x, y := <-c, <-c        
		fmt.Println("Đã nhận từ channel")
    fmt.Println(x, y, x+y)
}
