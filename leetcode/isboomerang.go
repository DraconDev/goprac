package leetcode

func isBoomerang(points [][]int) bool {

	if points[0][0] == points[1][0] && points[0][0] == points[2][0] {
		return false
	}

	return (points[0][1]-points[1][1])*(points[0][0]-points[2][0]) != (points[0][0]-points[1][0])*(points[0][1]-points[2][1])

}