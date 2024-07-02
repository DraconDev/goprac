package leetcode

func minOperations(nums []int, k int) int {

	arrayXor := 0
	// bitwise operations
	for _, item := range nums {
		arrayXor ^= item
	}

	arrayXor ^= k

	var result int
	for arrayXor != 0 {

		arrayXor &= (arrayXor - 1)
		result++
	}

	return result

}
