package main

import "fmt"

type Dog struct {
	gagau string
}
type Animal interface {
	Mount() string
}

func (t Dog) Mount() string {
	return "Cho co mom"
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func TestAssertion() {
	var Look any = "TrungLe"
	var ruby Animal = Dog{"Gau gau"}
	d, ok := ruby.(Dog)
	fmt.Println(d.Mount(), d.gagau, ok)

	do(32)
	do("Acquaman")
	test2, ok := Look.(int)
	fmt.Println("Check2", test2, ok)
}
