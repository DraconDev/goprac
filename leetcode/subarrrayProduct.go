package leetcode

// func numSubarrayProductLessThanK(nums []int, k int) int {
// 	result := 0

// 	for i := 0; i < len(nums); i++ {
// 		temp := 1
// 		for j := i; j < len(nums); j++ {
// 			temp *= nums[j]
// 			if temp < k {
// 				result++
// 			} else {
// 				break
// 			}
// 		}
// 	}

// 	return result

// }

func NumSubarrayProductLessThanK(nums []int, k int) int {
	var l, result int
	curr := 1
	for r := 0; r < len(nums); r++ {
		curr *= nums[r]
		for l <= r && curr >= k {
			curr /= nums[l]
			l++
		}
		result += r - l + 1
	}
	return result
}
