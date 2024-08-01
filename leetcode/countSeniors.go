package leetcode

func countSeniors(details []string) int {
	old := 0
	for _, v := range details {
		if v[11] >= '6' && v[12] >= '1' {
			old++
		} else if v[11] >= '7' {
			old++
		}
	}
	return old
}
