package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var hoten string
	fmt.Print("Nhap ho va ten cua ban: ")
	scanner := bufio.NewScanner((os.Stdin))
	if scanner.Scan() {
		hoten = scanner.Text()
	}
	fmt.Println("Xin chao", hoten)
	// TestGoroutine()
	// TestChannel()
	// TestBufferedChannel()
	// TestCloseChannel()
	// TestSelect()
	// Exercise()
	// TestMutex()
	// ExerciseMutex()
}
