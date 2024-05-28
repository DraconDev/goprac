package leetcode

func validPalindrome(s string) bool {

	if IsPalind(s) {
		return true
	}
	elem := s
	for i, _ := range s {
		elem = s[:i] + s[i+1:]
		if IsPalind(elem) {
			return true
		}
	}
	return false
}

func IsPalind(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
