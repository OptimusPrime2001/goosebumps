package main

import (
	"fmt"
	"math"
)

type Vertext struct {
	x, y float64
}
type Liver int

func (v Vertext) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (x Liver) jota() int {
	return int(x) + 2001
}
func TestMethod() {
	v := Vertext{3, 4}
	trung := Liver(20)
	fmt.Println(trung.jota())
	fmt.Println(v.abs())
}
