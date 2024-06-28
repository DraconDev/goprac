package leetcode

// func IsValidParans(check string) bool {
// 	if len(check)%2 == 1 || check == "" {
// 		return false
// 	}
// 	brackets := map[byte]byte{
// 		'(': ')',
// 		'[': ']',
// 		'{': '}',
// 	}
// 	stack := []byte{}
// 	for i := 0; i < len(check); i++ {
// 		for k, v := range brackets {
// 			if check[i] == k {
// 				stack = append(stack, check[i])
// 				break
// 			}
// 			if check[i] == v {
// 				if len(stack) == 0 || stack[len(stack)-1] != k {
// 					return false
// 				}
// 				stack = stack[:len(stack)-1]
// 				break
// 			}
// 		}
// 	}
// 	return len(stack) == 0
// }

func IsValidParans(check string) bool {
	if len(check)%2 == 1 || check == "" {
		return false
	}
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []rune{}
	for i := 0; i < len(check); i++ {
		for k, v := range brackets {
			if rune(check[i]) == k {
				stack = append(stack, rune(check[i]))
				break
			}
			if rune(check[i]) == v {
				if len(stack) == 0 || stack[len(stack)-1] != k {
					return false
				}
				stack = stack[:len(stack)-1]
				break
			}
		}
	}
	return len(stack) == 0
}
