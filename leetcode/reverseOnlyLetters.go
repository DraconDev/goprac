package leetcode

import "unicode"

func reverseOnlyLetters(s string) string {
	reversedLettersStack := []rune{}

	for _, v := range s {
		if unicode.IsLetter(v) {
			reversedLettersStack = append(reversedLettersStack, v)
		}
	}
	result := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			result = append(result, reversedLettersStack[len(reversedLettersStack)-1])
			reversedLettersStack = reversedLettersStack[:len(reversedLettersStack)-1]
		} else {
			result = append(result, v)
		}
	}
	return string(result)
}
