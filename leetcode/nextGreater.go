package leetcode

// func NextGreaterElement(nums1 []int, nums2 []int) []int {
// 	numbers := make(map[int]int)
// 	var max int
// 	for _, v := range nums2 {
// 		numbers[v]++
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	for _, v := range nums1 {
// 		numbers[v]--
// 	}

// 	var result []int
// 	for _, v := range nums1 {
// 		found := false

// 		for i := v + 1; i < max; i++ {
// 			if numbers[i] > 0 {
// 				result = append(result, i)
// 				numbers[i] = 0
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			result = append(result, -1)
// 		}
// 	}
// 	return result
// }

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	nums2Map := map[int]int{}

	for i, num := range nums2 {
		nums2Map[num] = i
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		found := false
		for nextIndex := nums2Map[num] + 1; nextIndex < len(nums2); nextIndex++ {
			if nums2[nextIndex] > num {
				result[i] = nums2[nextIndex]
				found = true
				break
			}
		}

		if !found {
			result[i] = -1
		}
	}

	return result
}
