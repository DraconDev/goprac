package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	// split to words
	allWords := strings.Split(s1+" "+s2, " ")

	// make map
	words := make(map[string]int)

	// add words to map
	for _, word := range allWords {
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
