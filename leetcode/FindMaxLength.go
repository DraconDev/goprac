package leetcode

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

func FindMaxLength(nums []int) int {
	res, sum := 0, 0
	prev := make(map[int]int)

	for i, val := range nums {
		if val == 1 {
			sum++
		} else {
			sum--
		}

		if sum == 0 {
			res = i + 1
		} else if prevId, ok := prev[sum]; !ok {
			prev[sum] = i
		} else {
			res = max(res, i-prevId)
		}
	}

	return res
}
