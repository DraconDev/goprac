package main

import "strings"

func bandNameGenerator(word string) string {
	if word == "" {
		return ""
	}
	lastIndex := len(word) - 1
	if word[0] == word[lastIndex] {
		return strings.ToUpper(string(word[0])) + word[1:lastIndex] + word
	}
	return "The " + strings.ToUpper(string(word[0])) + word[1:]
}
