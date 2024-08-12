package leetcode

func minDeletionSize(strs []string) int {
	minDel := 0

	if len(strs) == 1 {
		return minDel
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] > strs[j][i] {
				minDel++
				break
			}
		}
	}

	return minDel
}
