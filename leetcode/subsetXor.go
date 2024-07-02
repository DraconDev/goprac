package leetcode

func SubsetXORSum(nums []int) int {
	sumTotal := 0

	for _, num := range nums {
		sumTotal |= num
	}
	result := sumTotal << (len(nums) - 1)
	return result
}
