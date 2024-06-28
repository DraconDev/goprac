package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {

	s = strings.Trim(s, " ")
	// split string by space
	words := strings.Split(s, " ")

	return len(words[len(words)-1])

}
