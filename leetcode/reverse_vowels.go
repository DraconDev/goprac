package leetcode

import "strings"

func ReverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	result := []rune(s)
	l, r := 0, len(result)-1
	for l < r {
		if !strings.ContainsRune(vowels, result[l]) {
			l++
			continue
		}
		if !strings.ContainsRune(vowels, result[r]) {
			r--
			continue
		}
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}
	return string(result)
}
