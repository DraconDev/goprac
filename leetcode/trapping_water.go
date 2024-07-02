package leetcode

func Trapping_water(height []int) int {

	var left = 0
	var right = len(height) - 1
	var leftMax = 0
	var rightMax = 0
	var water = 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}

// func Trapping_water(height []int) int {

// 	var left = 0
// 	var right = len(height) - 1
// 	var leftMax = 0
// 	var rightMax = 0
// 	var water = 0
// 	for left < right {
// 		if height[left] < height[right] {
// 			if height[left] > leftMax {
// 				leftMax = height[left]
// 			} else {
// 				water += leftMax - height[left]
// 			}
// 			left++
// 		} else {
// 			if height[right] > rightMax {
// 				rightMax = height[right]
// 			} else {
// 				water += rightMax - height[right]
// 			}
// 		}
// 	}
// 	return water
// }

// func Trapping_water(height []int) int {
// 	var left = 0
// 	var right = len(height) - 1
// 	var maxLeft, maxRight = height[left], height[right]
// 	var trapLength = len(height)
// 	var area = 0
// 	var water = 0
// 	for left < right {
// 		if height[left] == 0 || height[left] < maxLeft {
// 			left++
// 			continue
// 		}
// 		if height[right] == 0 || height[right] < maxRight {
// 			right++
// 			continue
// 		}
// 	}
// 	return water
// }

// func Trapping_water(height []int) int {
// 	var left = 0
// 	var right = len(height) - 1
// 	var maxLeft, maxRight = height[left], height[right]
// 	min := 0
// 	water := 0
// 	for left < right {
// 		if height[left] == 0 {
// 			left++
// 			continue
// 		}
// 		if height[right] == 0 {
// 			right++
// 			continue
// 		}
// 		min = findMin(height[left], height[right])
// 		water += findWater(height[left+1:right-1], min)
// 		if height[left] < height[right] {
// 			left++
// 		} else {
// 			right--
// 		}

// 	}
// 	return water
// }

// func findMin(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func findWater(area []int, height int) int {
// 	total := sum(area...)
// 	return height*len(area) - total
// }

// func sum(numbers ...int) int {
// 	total := 0
// 	for _, num := range numbers {
// 		total += num
// 	}
// 	return total
// }
