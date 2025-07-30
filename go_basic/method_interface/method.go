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
func (u *Vertext) change(f float64) {
	u.x *= f
	u.y *= f
}
func otherChange(v *Vertext, f float64) {
	v.x *= f
	v.y *= f
}
func TestMethod() {
	v := &Vertext{3, 4}
	trung := Liver(20)
	fmt.Println(trung.jota())
	fmt.Println(v.abs())
	v.change(10)
	otherChange(v, 10)
	fmt.Println(v.abs())
}
