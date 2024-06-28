package leetcode

func MaxSubarrayLength(nums []int, k int) int {
	numbers := map[int]int{}

	var max int
	var sum int
	var left int
	var i int
	for i < len(nums) {
		if numbers[nums[i]] == k {
			numbers[nums[left]]--
			sum--
			left++
			continue
		}
		numbers[nums[i]]++
		sum++
		i++

		if sum > max {
			max = sum
		}
	}
	return max

}
