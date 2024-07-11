package leetcode

func ReverseParentheses(s string) string {

	stacks := [][]rune{}
	stacks = append(stacks, []rune{})
	params := 0
	for _, v := range s {
		if v == '(' {
			params++
			stacks = append(stacks, []rune{})
		} else if v == ')' {
			reverseElements(stacks[params])
			if params >= 1 {
				stacks[params-1] = append(stacks[params-1], stacks[params]...)
			}
			stacks[params] = []rune{}
			params--
		} else {
			stacks[params] = append(stacks[params], v)
		}
	}
	return string(stacks[0])
}

func reverseElements(s []rune) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
