package leetcode

func LargeGroupPositions(s string) [][]int {
	result := [][]int{}

	for i := 0; i < len(s)-2; i++ {
		count := 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				count++
				if j == len(s)-1 && count >= 3 {
					result = append(result, []int{i, j})
					i = j - 1
				}
			} else {
				if count >= 3 {
					result = append(result, []int{i, j - 1})
				}
				i = j - 1
				break
			}

		}

	}
	return result
}
