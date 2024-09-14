package leetcode

func balancedStringSplit(s string) int {
	count := 0
	right := 0
	for _, v := range s {
		if v == 'R' {
			right++
		} else {
			right--
		}
		if right == 0 {
			count++
		}
	}
	return count
}
