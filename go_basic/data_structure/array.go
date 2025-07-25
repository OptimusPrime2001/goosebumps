package main

import "fmt"

func TestArray() {
	var (
		students = [3]string{"Trung", "HaMi", "Hoat"}
		fordeer  = [...]int{3, 343}
	)
	fmt.Println(students, fordeer[1])
}
