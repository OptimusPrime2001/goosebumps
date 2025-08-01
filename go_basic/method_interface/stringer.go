package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (t Human) String() string {
	return fmt.Sprintf("%v (%v years old)", t.name, t.age)
}

func TestStringer() {
	dung := Human{
		"DungTran", 24,
	}
	trung := Human{"TrungLe", 32}

	fmt.Println(dung.String(), trung)
}
