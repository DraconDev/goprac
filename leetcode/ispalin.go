package leetcode

import "strconv"

func isPalin(x int) bool {
	for i := 0; i < len(strconv.Itoa(x)); i++ {
		if strconv.Itoa(x)[i] != strconv.Itoa(x)[len(strconv.Itoa(x))-1-i] {
			return false
		}
	}
	return true
}
