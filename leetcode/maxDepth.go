package leetcode

func maxDepth(s string) int {
	if s == "" {
		return 0
	}
	var params int
	maxNesting := 0

	for _, v := range s {
		if v == '(' {
			params++
		} else if v == ')' {
			params--
		}
		if params > maxNesting {
			maxNesting = params
		}
	}

	return maxNesting

}
