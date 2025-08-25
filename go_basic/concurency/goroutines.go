package main

import (
	"fmt"
	"time"
)
func say(s string) {
	for range 3 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func TestGoroutine() {
	go say("Hello Trung")
	say("Hello World")
}