package leetcode

func LongestOnes(nums []int, k int) int {
	i := 0
	j := 0
	for j < len(nums) {
		k -= 1 - nums[j]
		if k < 0 {
			k += 1 - nums[i]
			i++
		}
		j++
	}
	return j - i
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
