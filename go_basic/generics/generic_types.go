package main

import "fmt"


type Node[T any] struct {
	next *Node[T]
	val  T
}
func (l *Node[T]) Prepend(t T) *Node[T]{
	return &Node[T]{
		next:l,
		val:t,
	}
}
func (l *Node[T]) PrintAll() {
	for n:=l;n!=nil;n= n.next{
		fmt.Println(n.val)
	}
}
func TestGenericType[T any](t T)  {
	list := &Node[int]{}
	list = list.Prepend(20)
	list = list.Prepend(30)
	list.PrintAll()
}


