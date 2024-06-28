package leetcode

import "strconv"

func countPrimeSetBits(left int, right int) int {
	sets := 0
	for i := left; i <= right; i++ {
		bits := 0
		// convert i to binary
		binarrNum := strconv.FormatInt(int64(i), 2)
		for j := 0; j < len(binarrNum); j++ {
			if binarrNum[j] == '1' {
				bits++
			}
		}

		if bits == 2 || bits == 3 || bits == 5 || bits == 7 || bits == 11 || bits == 13 || bits == 17 || bits == 19 {
			sets++
		}
	}
	return sets
}
