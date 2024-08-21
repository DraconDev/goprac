package leetcode

import "strconv"

func bitwiseComplement(n int) int {

	// convert to binary
	binary := strconv.FormatInt(int64(n), 2)
	// flip bits
	converted := ""
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			converted += "1"
		} else {
			converted += "0"
		}
	}
	result, _ := strconv.ParseInt(converted, 2, 32)
	return int(result)
}
