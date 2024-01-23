package leetcode

func duplicates(s string) bool {
	var (
		m = make(map[rune]bool)
	)
	for _, r := range s {
		if m[r] {
			return true
		}
		m[r] = true
	}
	return false
}

func backtracking(arr []string, begging int, result string, skip []int, max int) int {
	for i := begging; i < len(arr); i++ {
		if skip[i] == 1 {
			continue
		}
		if !strings.ContainsAny(result, arr[i]) {
			skip[i] = 1
			if newMax := len(result) + len(arr[i]); newMax > max {
				max = newMax
			}
			max = backtracking(arr, i, result+arr[i], skip, max)
			skip[i] = 0
		}
	}
	return max
}

func MaxLength(arr []string) int {
	var (
		skip   = make([]int, len(arr))
		result string
		max    int
	)
	// Set duplicates strings to skip
	for i := 0; i < len(arr); i++ {
		if duplicates(arr[i]) {
			skip[i] = 1
		}
	}
	return backtracking(arr, 0, result, skip, max)
}