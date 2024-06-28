package leetcode

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxOperations(nums []int, k int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	operations := 0
	for val := range count {
		operations += minValue(count[val], count[k-val])
	}

	return operations / 2
}

// func MaxOperations(nums []int, k int) int {
// 	// * make hashmap of nums
// 	numMap := make(map[int][]int)
// 	for i, num := range nums {
// 		numMap[num] = append(numMap[num], i)
// 	}
// 	operations := 0

// 	for i := range numMap {
// 		if _, ok := numMap[k-i]; ok {
// 			if i*2 == k {
// 				operations += len(numMap[i]) / 2
// 				delete(numMap, i)
// 			} else {
// 				operations += minValue(len(numMap[i]), len(numMap[k-i]))
// 				delete(numMap, k-i)
// 				delete(numMap, i)

// 			}
// 		}
// 	}
// 	return operations
// }

// func maxOperations(nums []int, k int) int {
//     count := 0
//     seen := make(map[int]int)

//     for _, num := range nums {
//         res := k - num
//         if val, ok := seen[res]; ok && val > 0 {
//             count++
//             seen[res]--
//         } else {
//             seen[num]++
//         }
//     }

//     return count
// }

// func MaxOperations(nums []int, k int) int {
// 	// * make hashmap of nums
// 	numMap := make(map[int][]int)
// 	for i, num := range nums {
// 		numMap[num] = append(numMap[num], i)
// 	}
// 	operations := 0

// 	for i := range numMap {
// 		if _, ok := numMap[k-i]; ok {
// 			if i*2 == k {
// 				if len(numMap[k-i]) <= 1 {
// 					continue
// 				}
// 				nums[numMap[i][0]] = 0
// 				nums[numMap[k-i][1]] = 0
// 				operations += len(numMap[i]) / 2
// 				delete(numMap, i)
// 			} else {
// 				nums[numMap[i][0]] = 0
// 				nums[numMap[k-i][0]] = 0
// 				operations += minValue(len(numMap[i]), len(numMap[k-i]))
// 				delete(numMap, k-i)
// 				delete(numMap, i)

// 			}
// 		}
// 	}
// 	if operations == 0 {
// 		return 0
// 	}
// 	tempArr := []int{}
// 	for _, num := range nums {
// 		if num != 0 {
// 			tempArr = append(tempArr, num)
// 		}
// 	}
// 	nums = append(nums[:0], tempArr...)
// 	return operations
// }

// func MaxOperations(nums []int, k int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	operations := 0
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == k {
// 				operations++
// 				nums = append(nums[:i], nums[i+1:]...)
// 				nums = append(nums[:j-1], nums[j:]...)
// 				i--
// 				break
// 			}
// 		}
// 	}
// 	return operations
// }
