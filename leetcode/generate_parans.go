package leetcode

// func GenerateParenthesis(n int) []string {
// 	ans := make([]string, 0)
// 	current := make([]byte, n*2)
// 	backtrack(&ans, n, 0, 0, current)
// 	return ans
// }
// func backtrack(ans *[]string, n int, left int, right int, current []byte) {
// 	if left+right == n*2 {
// 		*ans = append(*ans, string(current))
// 		return
// 	}

// 	if left < n {
// 		current[left+right] = '('
// 		backtrack(ans, n, left+1, right, current)
// 	}

// 	if right < left {
// 		current[left+right] = ')'
// 		backtrack(ans, n, left, right+1, current)
// 	}
// }

func GenerateParenthesis(n int) []string {
	list := []string{}
	backtrack(&list, "", 0, 0, n)
	return list
}

func backtrack(list *[]string, str string, open int, close int, max int) {
	if len(str) == max*2 {
		*list = append(*list, str)
		return
	}

	if open < max {
		backtrack(list, str+"(", open+1, close, max)
	}
	if close < open {
		backtrack(list, str+")", open, close+1, max)
	}
}
