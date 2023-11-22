package leetcode

func LongestOnes(nums []int, k int) int {
	start := 0
	end := 0
	for end < len(nums) {
		k -= 1 - nums[end]
		if k < 0 {
			k += 1 - nums[start]
			start++
		}
		end++
	}
	return end - start
}

// func LongestOnes(nums []int, k int) int {
// 	max := 0
// 	currentMax := 0
// 	currentK := k
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 1 {
// 			currentMax++
// 		} else {
// 			if currentK > 0 {
// 				currentMax++
// 				currentK--
// 			} else {
// 				if currentMax > max {
// 					max = currentMax
// 				}
// 				currentMax = k
// 				// currentK = k
// 			}
// 		}
// 	}
// 	if max < currentMax {
// 		max = currentMax
// 	}
// 	return max
// }
