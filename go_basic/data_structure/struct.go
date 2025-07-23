package main

import "fmt"

func TestStruct() {
	type Liver struct {
		X int
		Y int
		Z string
	}

	type Staff struct {
		name   string
		age    int
		gender string
	}
	type Fordeer struct {
		headCount int
		staff     []Staff
	}
	listStaff := []Staff{{
		name:   "Trung Le",
		age:    20,
		gender: "Male",
	}, {
		name:   "HoatLa",
		age:    30,
		gender: "Female",
	}}

	var (
		allison = Liver{3, 4, "Don't know"}
		fordeer = Fordeer{20, listStaff}
	)
	allison.Z = "You never walk alone"

	p := &fordeer

	fmt.Println("TestStruct")
	fmt.Println(allison)
	fmt.Println(p.headCount, p.staff)
}
