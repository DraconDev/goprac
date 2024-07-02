package leetcode

import "strconv"

func SummaryRanges(nums []int) []string {
	var result []string

	if len(nums) == 0 {
		return result
	}

	start := ""
	for i, v := range nums[:len(nums)-1] {
		if v == nums[i+1]-1 {
			if start == "" {
				start = strconv.Itoa(v)
			}
			continue
		}
		if len(start) > 0 {
			result = append(result, start+"->"+strconv.Itoa(v))
			start = ""
		} else {
			result = append(result, strconv.Itoa(v))
			start = ""
		}

	}

	if len(start) > 0 {
		result = append(result, start+"->"+strconv.Itoa(nums[len(nums)-1]))
	} else {
		result = append(result, strconv.Itoa(nums[len(nums)-1]))
	}

	return result
}
