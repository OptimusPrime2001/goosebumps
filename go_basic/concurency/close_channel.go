package main

import "fmt"

func fibonacci(n int, c chan int){
	x,y := 0,1
	for range n{
		c<-x
		x,y = y,x+y
	}
	close(c)
}
func Test2(test chan int){
	test <- 1
	test <- 2
	// test <- 3
	fmt.Println("Gửi xong")
	close(test)
}
func TestCloseChannel() {
	
	c := make(chan int,3)
	test := make(chan int, 2)
	go Test2(test)
	go fibonacci(5,c)
	for v := range test {
		fmt.Println("Nhận:", v)
}

// Đọc hết dữ liệu từ channel c
for v := range c {
		fmt.Println("Fibonacci:", v)
}
	fmt.Println("Kết thúc")
}