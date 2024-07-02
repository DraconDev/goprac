package leetcode

import (
	"sort"
	"unicode"
)

func ShortestCompletingWord(licensePlate string, words []string) string {

	licensePlateLetters := make(map[rune]int)
	for _, v := range licensePlate {
		v = unicode.ToLower(v)
		if v >= 'a' && v <= 'z' {
			licensePlateLetters[v]++
		}
	}

	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

outer:
	for i, word := range words {
		letters := make(map[rune]int)
		for _, v := range word {
			v = unicode.ToLower(v)
			if v >= 'a' && v <= 'z' {
				letters[v]++
			}
		}
		for k, v := range licensePlateLetters {
			if letters[k] < v {
				continue outer
			}
		}
		return string(words[i])

	}

	return words[0]
}
