package leetcode

import "strings"

func toGoatLatin(sentence string) string {
	// spit the sentence into words
	words := strings.Split(sentence, " ")
	// create the result
	result := ""

	for i, word := range words {
		if strings.ContainsAny(string(word[0]), "aeiouAEIOU") {
			// if the word starts with a vowel
			result += word + "ma"
		} else {
			// if the word starts with a consonant
			result += word[1:] + string(word[0]) + "ma"
		}
		for j := 0; j < i+1; j++ {
			result += "a"
		}
		result += " "
	}
	return strings.Trim(result, " ")

}
