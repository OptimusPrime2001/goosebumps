package main

import "fmt"

type Human struct {
	name string
	age  int
}

func TestSlice() {
	student := []string{
		"Trung", "My", "Hoat",
	}
	friends := make([]string, 4)
	friends = append(friends, "Name", "Trung", "Darwin", "Robertson")
	newFriends := friends[4:6]

	workers := []Human{
		{
			name: "TrungLe",
			age:  22,
		}, {
			name: "HoangDung",
			age:  32,
		},
	}
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
		"Alex",
	}
	names[2] = "Herry"
	fmt.Println("======Test Slice======")
	fmt.Println("names worker", names, workers)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	q := []int{2, 3, 5, 7, 11, 13}
	b[0] = "XXX"
	fmt.Println("After change", a, b, q)
	fmt.Println(names)
	fmt.Println(student, len(friends), cap(friends), newFriends)
}
