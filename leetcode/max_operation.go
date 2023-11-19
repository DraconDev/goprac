package leetcode

// func MaxOperations(nums []int, k int) int {
// 	// * make hashmap of nums
// 	numMap := make(map[int][]int)
// 	for i, num := range nums {
// 		numMap[num] = append(numMap[num], i)
// 	}
// 	operations := 0
// 	for i, v := range nums {
// 		if v == 0 {
// 			continue
// 		}
// 		if _, ok := numMap[k-v]; ok {
// 			if v*2 == k && len(numMap[k-v]) > 1 {
// 				operations++
// 				nums[numMap[k-v][0]] = 0
// 				nums[numMap[k-v][1]] = 0
// 				nums[numMap[k-v]] = append(nums[numMap[k-v]][:2], nums[numMap[k-v]][2:]...)

// 			} else {
// 				operations++
// 				nums[numMap[k-v][0]] = 0
// 				nums[i] = 0
// 			}
// 		}
// 	}

// 	filteredNums := nums[:0] // create a new slice with length 0

// 	for _, num := range nums {
// 		if num != 0 {
// 			filteredNums = append(filteredNums, num)
// 		}
// 	}

// 	nums = append([]int{}, filteredNums...)

// 	return operations
// }

func MaxOperations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	operations := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == k {
				operations++
				nums = append(nums[:i], nums[i+1:]...)
				nums = append(nums[:j-1], nums[j:]...)
				i--
				break
			}
		}
	}
	return operations
}
