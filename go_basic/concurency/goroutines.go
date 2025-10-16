package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d is running\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d is done\n", id)
}
func TestGoroutine() {
	start := time.Now()
	task(1)
	task(2)
	task(3)
	fmt.Printf("Total time: %v\n", time.Since(start))
}
