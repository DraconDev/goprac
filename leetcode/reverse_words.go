package leetcode

import "strings"

func ReverseWords(s string) string {
	trimmedStr := strings.TrimSpace(s)
	result := strings.Split(trimmedStr, " ")
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return strings.Join(result, " ")
}