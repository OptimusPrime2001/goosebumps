package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result_slice := make([][]uint8, dy)
	for x := range result_slice {
		ele_slice := make([]uint8, dx)
		for y := range ele_slice {
			ele_slice[y] = uint8((x + y) / 2)
		}
		result_slice[x] = ele_slice
	}
	return result_slice
}
func ExerciseSlices() {
	pic.Show(Pic)
}
