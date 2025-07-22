package main

import "fmt"

func Variable() {
	var fullname string = "Trung Le"
	var age int = 24

	var score int
	var isActive bool

	lover := "no thing"
	count := 10
	active := true

	var x, y, z int
	a, b := 10, "Go lang"

	var (
		name     string  = "Go lang"
		version  float32 = 1.32
		isStable bool    = true
	)
	fmt.Print(fullname, age, score, isActive, lover, count, active, x, y, z, a, b, name, version, isStable)

	const Pi = 3.14
	const Greeting = "Hello"
}
