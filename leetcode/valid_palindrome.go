package leetcode

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	pattern := "[^a-zA-Z0-9]"
	s = strings.ToLower(regexp.MustCompile(pattern).ReplaceAllString(s, ""))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
