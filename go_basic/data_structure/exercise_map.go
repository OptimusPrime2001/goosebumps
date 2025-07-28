package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	split_s := strings.Fields(s)
	result := make(map[string]int)
	for _, value := range split_s {
		el, ok := result[value]
		if ok {
			result[value] = el + 1
		} else {
			result[value] = 1
		}
	}
	return result
}

func ExerciseMaps() {
	wc.Test(WordCount)
}
