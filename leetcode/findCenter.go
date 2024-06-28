package leetcode

func findCenter(edges [][]int) int {
	points := make(map[int]int)
	for _, edge := range edges {
		points[edge[0]]++
		points[edge[1]]++
		if points[edge[0]] > 1 {
			return edge[0]
		}
		if points[edge[1]] > 1 {
			return edge[1]
		}
	}

	return 0

}
