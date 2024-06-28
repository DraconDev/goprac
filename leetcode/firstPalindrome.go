package leetcode

func firstPalindrome(words []string) string {

	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(word string) bool {
	length := len(word)
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-i-1] {
			return false
		}
	}
	return true
}
