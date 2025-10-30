package main

import (
	"fmt"
	"sync"
)

func TestMutextExample() {
	tokens := 0
	wg := sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go func() {
			tokens++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("tokens:", tokens)
}

// Mỗi goroutine tăng counter cùng lúc, không có đồng bộ hóa.
// Toán tử counter++ thực chất gồm 3 bước riêng biệt:
// Đọc giá trị hiện tại của counter
// Cộng thêm 1
// Ghi lại vào bộ nhớ
// Vì vậy, nếu hai goroutine cùng đọc rồi ghi gần như cùng lúc, một trong hai phép tăng sẽ bị mất → race condition.
