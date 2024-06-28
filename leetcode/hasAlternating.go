package leetcode

import "strconv"

func hasAlternatingBits(n int) bool {
	// convert to binary
	binary := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(binary)-1; i++ {
		if binary[i] == binary[i+1] {
			return false
		}
	}
	return true
}
