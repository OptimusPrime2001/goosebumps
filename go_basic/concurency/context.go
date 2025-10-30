package main

import (
	"context"
	"fmt"
	"time"
)

type priorityKey string

const keyPriority priorityKey = "priority"

func employee(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Employee context canceled")
			return
		default:
			priority := ctx.Value(keyPriority)
			fmt.Println("Employee with priority:", priority)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func TestContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, keyPriority, "high")
	go employee(ctx)
	time.Sleep(3 * time.Second)
}
