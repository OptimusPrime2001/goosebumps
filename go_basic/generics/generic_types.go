package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func (l *Node[T]) Prepend(t T) *Node[T] {
	return &Node[T]{
		next: l,
		val:  t,
	}
}
func (l *Node[T]) PrintAll() {
	for n := l; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}
func TestGenericType[S any](s S) {
	fmt.Printf("Type of s is %T\n", s)
	list := &Node[S]{}

	list = list.Prepend(s)
	list = list.Prepend(any(30).(S))
	list.PrintAll()
}
