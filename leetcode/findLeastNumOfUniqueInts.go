package leetcode

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	var grouped = make(map[int]int)

	for _, v := range arr {
		grouped[v]++
	}

	count := make([]int, len(arr)+1)

	for _, v := range grouped {
		count[v] += 1
	}

	for i := 1; i < len(count); i++ {
		for k > 0 && count[i] > 0 {
			if k >= i {
				count[i]--
			}
			k -= i

		}
		if k <= 0 {
			break
		}

	}

	var unique int

	for _, v := range count[1:] {
		if v > 0 {
			unique += v
		}
	}

	return unique
}
