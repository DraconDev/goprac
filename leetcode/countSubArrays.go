package leetcode

func CountSubarrays(nums []int, k int) int64 {

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	idx := []int{} // index of maximum element
	ret := int64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == max {
			// keep len(idx) <= k
			if len(idx) == k {
				idx = idx[1:]
			}
			idx = append(idx, i)
		}
		if len(idx) == k {
			// calculate how many elements between 0 and idx[0]
			// ret += idx[0] - 0 + 1
			ret += int64(idx[0]) + 1
		}
	}
	return ret
}
