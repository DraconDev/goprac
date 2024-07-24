package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	// split to words
	s1Words := strings.Split(s1, " ")
	s2Words := strings.Split(s2, " ")
	// make map
	words := make(map[string]int)

	// add words to map
	for _, word := range s1Words {
		words[word]++
	}
	for _, word := range s2Words {
		words[word]++
	}
	// find uncommon words
	var uncommonWords []string
	for word := range words {
		if words[word] == 1 {
			uncommonWords = append(uncommonWords, word)
		}
	}

	return uncommonWords
}
