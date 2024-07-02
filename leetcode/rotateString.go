package leetcode

import "strings"

func RotateString(s string, goal string) bool {

	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	// find string in goal
	goal = goal + goal
	// find word in string

	if strings.Contains(goal, s) {
		return true
	}
	return false
}
