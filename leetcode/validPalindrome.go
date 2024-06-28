package leetcode

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalind(s, left+1, right) || isPalind(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

func isPalind(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
