package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := range int(100e8) {
		sum += i
	}
}
func TestConcuDemo() {
	numCPU := 1
	fmt.Println("numCPU:", numCPU)
	runtime.GOMAXPROCS(numCPU)
	start := time.Now()
	wg := sync.WaitGroup{}
	for range 10 {
		wg.Add(1)
		go heavyTask(&wg)
	}
	wg.Wait()
	fmt.Println("heavyTask took:", time.Since(start))
}
