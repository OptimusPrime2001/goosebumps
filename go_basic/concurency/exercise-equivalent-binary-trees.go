package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}
func walkDone(fn func(), ch chan int) {
	fn()
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walk(t1, ch1)
	Walk(t2, ch2)
	for range ch1 {
		v1, v2 := <-ch1, <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func Exercise() {
	ch := make(chan int)
	go walkDone(func() {
		Walk(tree.New(3), ch)
	}, ch)
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(Same(tree.New(3), tree.New(3)))
	fmt.Println(Same(tree.New(3), tree.New(4)))
}
