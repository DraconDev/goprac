package leetcode

import "strconv"

func binaryGap(n int) int {

	max := 0
	gap := 1
	first := true
	binaryInString := []rune(strconv.FormatInt(int64(n), 2))
	for i := 0; i < len(binaryInString); i++ {
		if binaryInString[i] == '0' {
			gap++
		} else {
			if first {
				first = false
				gap = 1
				continue
			}
			if gap > max {
				max = gap
			}
			gap = 1
		}
	}

	if max == 0 {
		return max
	}

	return max
}
