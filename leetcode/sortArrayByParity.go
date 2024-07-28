package leetcode

// func sortArrayByParity(nums []int) []int {
// 	odds := []int{}
// 	evens := []int{}
// 	for _, v := range nums {
// 		if v%2 == 0 {
// 			evens = append(evens, v)
// 		} else {
// 			odds = append(odds, v)
// 		}
// 	}
// 	return append(evens, odds...)

// }

func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 == 1 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums

}
