package leetcode

import "sort"

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}

	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i][1] == coordinates[j][1] {
			return coordinates[i][1] < coordinates[j][1]
		}
		return coordinates[i][1] < coordinates[j][1]
	})

	if horizontal(coordinates) {
		return true
	}

	slope, ok := coors(coordinates[0], coordinates[1])
	if !ok {
		return false
	}

	for i := 1; i < len(coordinates)-1; i++ {
		currentSlope, ok := coors(coordinates[i], coordinates[i+1])
		if !ok || currentSlope != slope {
			return false
		}
	}
	return true
}

func coors(a, b []int) (float64, bool) {
	if a[1] == b[1] {
		return 0, false
	}
	x := (b[1] - a[1])
	y := (b[0] - a[0])
	slope := float64(y) / float64(x)
	return slope, true

}

func horizontal(cords [][]int) bool {
	return cords[0][1] == cords[len(cords)-1][1]
}
