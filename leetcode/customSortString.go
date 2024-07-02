package leetcode

import "strings"

func customSortString(order string, s string) string {

	letters := make(map[string]int)

	for _, c := range s {
		letters[string(c)]++
	}

	var result string

	for _, c := range order {
		if letters[string(c)] > 0 {
			result += strings.Repeat(string(c), letters[string(c)])
			delete(letters, string(c))

		}
	}

	for k, v := range letters {
		for i := 0; i < v; i++ {
			result += k
		}
	}

	return result

}
