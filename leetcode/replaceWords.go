package leetcode

import (
	"sort"
	"strings"
)

func ReplaceWords(dictionary []string, sentence string) string {
	// sort dictionary
	sort.Strings(dictionary)
	// array of words
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for _, dict := range dictionary {
			if strings.HasPrefix(word, dict) {
				words[i] = dict
				break
			}
		}
	}
	return strings.Join(words, " ")

}
