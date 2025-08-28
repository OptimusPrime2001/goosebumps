package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) deposit(amout int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amout
}
func (b *BankAccount) withdraw(amout int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.balance > amout {
		b.balance -= amout
		fmt.Printf("Số tiền còn lại sau khi rút là: %d\n", b.balance)
	}
	if b.balance < amout {
		fmt.Println("Bạn không còn đủ tiền để rút")
	}
}

// Lấy số dư
func (a *BankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func ExerciseMutex() {
	account := BankAccount{
		balance: 200_000_000,
	}
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			account.deposit(1_000_000 * (i + 1))
		}(i)
	}
	// Tạo 5 goroutine rút tiền
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			account.withdraw(2_000_000 * (i + 1))
		}(i)
	}
	wg.Wait()
	fmt.Println("Số dư cuối cùng:", account.Balance())

	// Chờ để log hiển thị đẹp hơn
	time.Sleep(time.Second)

}
