package main

import (
	"strings"
)

func bandNameGenerator(word string) string {
	// caser := cases.Title(language.English)
	if word == "" {
		return ""
	}
	lastIndex := len(word) - 1
	equal := strings.EqualFold(string(word[0]), string(word[lastIndex]))
	if equal {
		return strings.ToUpper(string(word[0])) + word[1:lastIndex] + word
	}

	// return "The " + caser.String(word)
	return "The " + strings.Title(string(word))
}
