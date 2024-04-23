package leetcode

import "strings"

func detectCapitalUse(word string) bool {

	lower := strings.ToLower(word)
	higher := strings.ToUpper(word)
	capital := strings.Title(strings.ToLower(word))
	if word == lower || word == higher || word == capital {
		return true
	}
	return false
}
