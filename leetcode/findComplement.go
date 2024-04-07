package leetcode

import "strconv"

func findComplement(num int) int {
	// nums as bit string
	numStr := strconv.FormatInt(int64(num), 2)

	// complement as bit string
	complementStr := ""
	for _, v := range numStr {
		if string(v) == "0" {
			complementStr += "1"
		} else {
			complementStr += "0"
		}
	}
	// convert to int
	complement, _ := strconv.ParseInt(complementStr, 2, 64)
	return int(complement)
}
