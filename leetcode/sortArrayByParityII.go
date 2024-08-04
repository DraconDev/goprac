package leetcode

func sortArrayByParityII(nums []int) []int {
	odd := 1
	even := 0
	result := make([]int, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			result[even] = v
			even += 2
		} else {
			result[odd] = v
			odd += 2
		}
	}
	return result
}
