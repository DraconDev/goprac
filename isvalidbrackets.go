package main

func isValidBrackets(check string) bool {
	runes := []rune(check)
	both := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	brackets := []rune{}
	for i := 0; i < len(runes); i++ {
		for k, v := range both {
			if runes[i] == k {
				brackets = append(brackets, runes[i])
				break
			}
			if runes[i] == v {
				if len(brackets) == 0 || brackets[len(brackets)-1] != k {
					return false
				}
				brackets = brackets[:len(brackets)-1]
				break
			}
		}
	}
	return len(brackets) == 0
}
