package main

import "fmt"

func TestSlice() {
	student := []string{
		"Trung", "My", "Hoat",
	}
	friends := make([]string, 4)
	friends = append(friends, "Name", "Trung", "Darwin", "Robertson")
	newFriends := friends[4:6]

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("======Test Slice======")
	fmt.Println("names", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println("After change", a, b)
	fmt.Println(names)
	fmt.Println(student, len(friends), cap(friends), newFriends)
}
