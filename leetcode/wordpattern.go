package leetcode

import "strings"

func WordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	dictionary := make(map[rune]string)
	used := make(map[string]bool)

	for i, c := range pattern {
		if dictionary[c] == "" {
			if used[words[i]] {
				return false
			}
			used[words[i]] = true
			dictionary[c] = words[i]
		} else {
			if dictionary[c] != words[i] {
				return false
			}
		}
	}

	return true

}
