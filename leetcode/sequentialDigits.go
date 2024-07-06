package leetcode

import "strconv"

func SequentialDigits(low int, high int) []int {
	var result []int
	digits := "0123456789"
	length := len(strconv.Itoa(low))
	i := int(strconv.Itoa(low)[0] - '0')

	for len(result) == 0 || result[len(result)-1] <= high {

		if length >= 10 {
			break
		}

		if i+length > 10 {
			length++
			i = 1
			continue
		}
		add, _ := strconv.Atoi(digits[i : i+length])

		i++
		if add < low {
			continue
		}
		result = append(result, add)
	}

	if length >= 10 {
		return result
	}

	return result[:len(result)-1]

}
