package main

import "math"

func CheckForFactor(base int, factor int) bool {
	return math.Mod(float64(base), float64(factor)) == 0
	// your code here
}
