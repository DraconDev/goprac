package leetcode

func MoveZeroes(nums []int) {
	result := []int{}

	for _, v := range nums {
		if v != 0 {
			result = append(result, v)
		}
	}
	for i := range nums {
		if i < len(result) {
			nums[i] = result[i]
		} else {
			nums[i] = 0
		}
	}
}
