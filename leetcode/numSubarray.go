package leetcode

// func NumSubarraysWithSum(nums []int, goal int) int {

// 	var sum int
// 	var curSum int

// 	for i := range nums {
// 		for _, v2 := range nums[i:] {
// 			curSum += v2
// 			if curSum == goal {
// 				sum++
// 			} else if curSum > goal {
// 				break
// 			}
// 		}
// 		curSum = 0
// 	}

// 	return sum

// }

func sumAtMost(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}

	res, prev, cur := 0, 0, 0
	for i, val := range nums {
		cur += val
		for cur > goal {
			cur -= nums[prev]
			prev++
		}
		res += i - prev + 1
	}

	return res
}

func NumSubarraysWithSum(nums []int, goal int) int {
	return sumAtMost(nums, goal) - sumAtMost(nums, goal-1)
}
