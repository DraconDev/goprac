package leetcode

func LongestCommonSubsequence(text1 string, text2 string) int {
	var first_indexes = make(map[rune][]int)

	for i, r := range text1 {
		first_indexes[r] = append(first_indexes[r], i)
	}

	var max_sub int
	// go through all letters
	count := 0
	var last int
	for i, r := range text2 {
		if indexes, ok := first_indexes[r]; ok && indexes[len(indexes)-1] >= last {
			count++
			for i2 := range indexes {
				if indexes[i2] >= last {
					last = indexes[i2]
					break
				}
			}
		} else {
			count = 0
			last = i
			continue
		}
		if count > max_sub {
			max_sub = count
		}

	}

	return max_sub

}
