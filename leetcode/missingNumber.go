package leetcode

func missingNumber(nums []int) int {
	// 	sort.Ints(nums)

	// 	for i, v := range nums {
	// 		if v != i {
	// 			return i
	// 		}
	// 	}
	// 	return nums[len(nums)-1] + 1
	// }

	dict := make([]bool, len(nums)+1)

	for _, n := range nums {
		dict[n] = true
	}

	for i, exists := range dict {
		if !exists {
			return i
		}
	}

	return 0
}
