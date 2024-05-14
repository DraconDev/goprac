package leetcode

func MaxCount(m int, n int, ops [][]int) int {
	result := [][]int{}
	for range m {
		row := make([]int, n)
		result = append(result, row)
	}

	for i := 0; i < len(ops); i++ {
		for j := 0; j < ops[i][0]; j++ {
			for k := 0; k < ops[i][1]; k++ {
				result[j][k]++
			}
		}
	}
	max := result[0][0]
	count := 0
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			if result[i][j] == max {
				count++
			}
		}
	}
	return count
}
