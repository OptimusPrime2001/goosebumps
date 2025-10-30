package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age  int
}

func TestSlice() {
	student := []string{
		"Trung", "My", "Hoat",
	}
	Tung := [1]Human{
		{
			name: "TrungLe",
			age:  22,
		},
	}
	fmt.Println("Tung co phai la slice hay khong", reflect.TypeOf(Tung).Kind() == reflect.Slice)
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

	str1 := []string{"Apple1", "Apple2", "Apple3"}
	str2 := []string{"Apple4", "Apple5", "Apple6"}
	str3 := []string{"Apple7", "Apple8", "Apple9"}
	str4 := append(str1, str2...)
	str4 = append(str4, str3...+)
	fmt.Println(str4)
}
