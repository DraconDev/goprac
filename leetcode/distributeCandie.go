package leetcode

func distributeCandie(candies int, num_people int) []int {
	result := make([]int, num_people)

	giveAmount := 1
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			result[i] += min(candies, giveAmount)
			candies -= giveAmount
			giveAmount++
			if candies <= 0 {
				return result
			}
		}
	}
	return result

}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
