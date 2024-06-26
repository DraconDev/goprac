package leetcode

func MaximumStrongPairXor(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			x := nums[i]
			y := nums[j]
			if abs(y-x) <= min(y, x) {
				res = max(res, x^y)
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
