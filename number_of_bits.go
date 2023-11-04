package main

import "strconv"

// func count_bits(n int) int {
// 	result := 0
// 	for _, v := range strconv.Itoa(n) {
// 		if v == '1' {
// 			result++
// 		}
// 	}
// 	return result
// }

func hammingWeight(num uint32) int {
	result := 0
	converted := strconv.FormatUint(uint64(num), 2)
	for _, v := range converted {
		if v == '1' {
			result++
		}
	}
	return result
}
