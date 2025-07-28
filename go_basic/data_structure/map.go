package main

import "fmt"

type Dev struct {
	name string
	age  int
}

var newMap map[string]Dev

func TestMap() {
	newMap = make(map[string]Dev)
	newMap["DevFrontend"] = Dev{
		name: "TrungLe",
		age:  22,
	}
	newMap["DevBackend"] = Dev{
		"HungDinh",
		23,
	}
	fmt.Println(newMap)
}
