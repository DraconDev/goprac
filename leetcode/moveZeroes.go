package leetcode

// func MoveZeroes(nums []int) {
// 	result := []int{}

// 	for _, v := range nums {
// 		if v != 0 {
// 			result = append(result, v)
// 		}
// 	}
// 	for i := range nums {
// 		if i < len(result) {
// 			nums[i] = result[i]
// 		} else {
// 			nums[i] = 0
// 		}
// 	}
// }

// func MoveZeroes(nums []int) []int {
// 	zeroIndex := 0
// 	for i, v := range nums {
// 		if v != 0 {
// 			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
// 			zeroIndex++
// 		}
// 		if i > len(nums)-zeroIndex {
// 			break
// 		}
// 	}
// 	return nums
// }

// func MoveZeroes(nums []int) []int {
// 	zeroIndex := 0
// 	for i, v := range nums {
// 		if v != 0 {
// 			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
// 			zeroIndex++
// 		}
// 	}
// 	return nums
// }

func MoveZeroes(nums []int) {
	snowBallSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowBallSize++
		} else if snowBallSize > 0 {
			nums[i], nums[i-snowBallSize] = nums[i-snowBallSize], nums[i]
		}
	}
}
