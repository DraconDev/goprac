package main

func findNumber(nums []int, target int) int {
	
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// func findElement(nums []int, target int) int {
// 	var l, r = 0, len(nums) - 1

// 	for l <= r {
// 		middle := (l + r) / 2
// 		if nums[middle] == target {
// 			return middle
// 		}
// 		if nums[middle] < target {
// 			l = middle + 1
// 		} else {
// 			r = middle - 1
// 		}
// 	}
// 	return -1
// }
