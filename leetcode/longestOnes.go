package leetcode

func LongestOnes(nums []int, k int) int {
	front := 0
	back := 0
	for back < len(nums) {
		if nums[back] == 0 {
			k--
		}
		if k < 0 {
			if nums[front] == 0 {
				k++
			}
			front++
		}
		back++
	}
	return back - front
}

//
