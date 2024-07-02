package leetcode

import "strings"

func findWords(words []string) []string {
	rows := []map[string]bool{
		{"Q": true, "W": true, "E": true, "R": true, "T": true, "Y": true, "U": true, "I": true, "O": true, "P": true},
		{"A": true, "S": true, "D": true, "F": true, "G": true, "H": true, "J": true, "K": true, "L": true},
		{"Z": true, "X": true, "C": true, "V": true, "B": true, "N": true, "M": true},
	}

	result := []string{}

	for _, word := range words {
		isRow := true

		for _, row := range rows {
			isRow = true
			for _, letter := range word {
				if !row[strings.ToUpper(string(letter))] {
					isRow = false
					break
				}
			}
			if isRow == true {
				break
			}
		}

		if isRow {
			result = append(result, word)
		}
	}
	return result
}
