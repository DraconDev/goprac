package leetcode

func singleNumber(nums []int) []int {
	elems := make(map[int]int)
	for _, v := range nums {
		elems[v]++
	}
	result := make([]int, 0)
	for k, v := range elems {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
