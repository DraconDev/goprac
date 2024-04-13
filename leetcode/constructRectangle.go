package leetcode

import "math"

func ConstructRectangle(area int) []int {
	// square root area
	width := int(math.Sqrt(float64(area)))
	height := int(math.Sqrt(float64(area)))

	for width*height != area {
		if width*height < area {
			width++
		} else {
			height--
		}
	}
	return []int{width, height}
}
