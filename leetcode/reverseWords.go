package leetcode

import "strings"

func reverseWords(s string) string {
	var words []string
	// split by space
	words = strings.Split(s, " ")

	for word := range words {
		// reverse string of each word
		words[word] = reverseWord(words[word])
	}

	// join all words
	return strings.Join(words, " ")
}

func reverseWord(word string) string {
	// convert to array
	wordArr := []byte(word)

	for i := 0; i < len(wordArr)/2; i++ {
		wordArr[i], wordArr[len(wordArr)-i-1] = wordArr[len(wordArr)-i-1], wordArr[i]
	}
	return string(wordArr)
}
