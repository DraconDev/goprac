package leetcode

import "strconv"

func CountBits(n int) []int {

	var result []int

	for i := 0; i <= n; i++ {
		// convert int to binry
		binary := strconv.FormatInt(int64(i), 2)
		count := 0
		for j := 0; j < len(binary); j++ {
			if binary[j] == '1' {
				count++
			}
		}
		result = append(result, count)
	}

	return result
}
