package leetcode

func diStringMatch(s string) []int {

	var result []int
	min, max := 0, len(s)
	for _, v := range s {
		if v == 'I' {
			result = append(result, min)
			min++
		} else {
			result = append(result, max)
			max--
		}
	}
	result = append(result, min)
	return result

}
