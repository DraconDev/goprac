package leetcode

import (
	"fmt"
)

func MaximumStrongPairXor(nums []int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i]*2 || nums[i] > nums[j]*2 {
				continue
			}
			fmt.Printf("%b %b\n", nums[i], nums[j])
			difference := nums[i] ^ nums[j]
			if difference > result {
				result = difference
			}
		}
	}

	return result

}
