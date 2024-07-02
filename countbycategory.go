package main

import (
	"unicode"
)

func Solve(s string) []int {
	//..
	counter := make(map[string]int)
	for _, v := range s {
		if unicode.IsUpper(v) {
			counter["upper"]++
		} else if unicode.IsLower(v) {
			counter["lower"]++
		} else if unicode.IsDigit(v) {
			counter["digit"]++
		} else {
			counter["other"]++
		}

	}
	result := []int{counter["upper"], counter["lower"], counter["digit"], counter["other"]}
	return result
}
