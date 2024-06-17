package leetcode

import "math"

func JudgeSquareSum(c int) bool {

	squares := make(map[int]bool)

	for i := int(math.Sqrt(float64(c))); i >= 0; i-- {
		squares[i*i] = true
	}

	for k := range squares {
		if squares[c-k] {
			return true
		}
	}
	return false

}
