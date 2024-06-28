package leetcode

import "strconv"

func AddDigits(num int) int {

	converted := strconv.Itoa(num)
	for len(converted) > 1 {
		sum := 0
		for _, v := range converted {
			digit, _ := strconv.Atoi(string(v))
			sum += digit
		}
		converted = strconv.Itoa(sum)
	}
	convertedInt, _ := strconv.Atoi(converted)
	return convertedInt
}
