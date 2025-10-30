package main

import (
	"context"
	"fmt"
	"time"
)

func cookPho(ctx context.Context, ch chan string) {
	fmt.Println("Start Cooking Pho")
	select {
	case <-time.After(1 * time.Second):
		ch <- "Cook Pho Done"
		return
	case <-ctx.Done():
		fmt.Println("Cook Pho context canceled")
		return
	}
}
func cookRice(ctx context.Context, ch chan string) {
	fmt.Println("Start Cooking Rice")
	select {
	case <-time.After(2 * time.Second):
		ch <- "Cook Rice Done"
		return
	case <-ctx.Done():
		fmt.Println("Cook Rice context canceled")
		return
	}
}
func TestContextExample() {
	chPho := make(chan string)
	chRice := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	go cookPho(ctx, chPho)
	go cookRice(ctx, chRice)
	for range 2 {
		time.Sleep(2 * time.Second)
		select {
		case msg := <-chPho:
			fmt.Println("msg from chPho:", msg)
		case msg := <-chRice:
			fmt.Println("msg from chRice:", msg)
		case <-ctx.Done():
			fmt.Println("Context canceled")
		default:
			fmt.Println("No message yet")
		}
	}
}
