package leetcode

func duplicateZeros(arr []int) {
	result := []int{}
	for _, v := range arr {
		if v == 0 {
			result = append(result, 0)
			result = append(result, 0)
		} else {
			result = append(result, v)
		}
	}
	copy(arr, result)
}
