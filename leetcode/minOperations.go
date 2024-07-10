package leetcode

func minOperations(logs []string) int {
	count := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if count > 0 {
				count--
			}
		} else if logs[i] != "./" {
			count++
		}
	}
	return count

}
