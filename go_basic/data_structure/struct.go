package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Staff struct {
	TEN      string `json:"name"`
	TUOI     int    `json:"age"`
	GIOITINH string `json:"gender"`
}
type Fordeer struct {
	headCount int
	staff     []Staff
}
type Person struct {
	age    int
	life   int
	status string
}

func NewPerson(age int) *Person {
	return &Person{
		age:    age,
		life:   100,
		status: "Still alive",
	}
}

// Value receiver
func (nv Staff) getInforStaff() {
	fmt.Println(nv.TEN, nv.TUOI, nv.GIOITINH)
}

// Pointer receiver
func (f *Person) killPerson() {
	f.status = "You are dead"
	f.life = 0
}
func TestStruct() {

	listStaff := []Staff{{
		TEN:      "Trung Le",
		TUOI:     20,
		GIOITINH: "Male",
	}, {
		TEN:      "HoatLa",
		TUOI:     30,
		GIOITINH: "Female",
	}}
	for _, nv := range listStaff {
		nv.getInforStaff()
	}
	var (
		allison = Person{3, 4, "Still alive"}

		fordeer = Fordeer{20, listStaff}
	)
	allison.status = "To young"
	allison.killPerson()
	p := &fordeer
	fmt.Println("allison:", allison)
	fmt.Println("fordeer:", p)

	output, err := json.Marshal(listStaff)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(output))
}
