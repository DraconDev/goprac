package leetcode

func IncreasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}