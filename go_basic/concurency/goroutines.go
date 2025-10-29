package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d is running\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d is done\n", id)
}
func runSequential() {
	// Hàm mô tả các task chạy tuần tự trong go
	start := time.Now()
	task(1)
	task(2)
	task(3)
	fmt.Printf("Total time: %v\n", time.Since(start))
	// Thời gian cần để hoàn thành 3 task sẽ mất 3 giây
}
func runConcurrent() {
	// Hàm mô tả các task chạy song song trong go
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		task(1)
	}()
	go func() {
		defer wg.Done()
		task(2)
	}()
	go func() {
		defer wg.Done()
		task(3)
	}()
	wg.Wait()
	go task(1)
	go task(2)
	go task(3)
	fmt.Printf("Total time: %v\n", time.Since(start))
	// Thời gian cần để hoàn thành 3 task sẽ mất 1 giây
}
func TestGoroutine() {
	// runSequential()

	fmt.Println("=======>Sequential tasks <=======")
	runSequential()

	fmt.Println("=======>Concurrent tasks <=======")
	runConcurrent()
}
