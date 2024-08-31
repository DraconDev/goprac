package leetcode

import "strings"

func findOcurrences(text string, first string, second string) []string {

	textArr := strings.Split(text, " ")
	var result []string

	for i := 0; i < len(textArr)-2; i++ {
		if textArr[i] == first && textArr[i+1] == second {
			result = append(result, textArr[i+2])
		}
	}
	return result

}
