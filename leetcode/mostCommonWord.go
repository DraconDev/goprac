package leetcode

import (
	"strings"
	"unicode"
)

func MostCommonWord(paragraph string, banned []string) string {
	// split up paragraph into words map
	paragraph = strings.ToLower(paragraph)

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		// split on non letter
		return !unicode.IsLetter(r)
	})

	wordsMap := map[string]int{}
	for _, word := range words {
		wordsMap[word]++
	}

	bannedWords := map[string]bool{}

	for _, word := range banned {
		bannedWords[word] = true
	}

	// make struct highest

	highest := struct {
		word  string
		count int
	}{}

	for word, count := range wordsMap {
		if bannedWords[word] {
			continue
		}
		if count > highest.count {
			highest.word = word
			highest.count = count
		}
	}

	return highest.word
}
