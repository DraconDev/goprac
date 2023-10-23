package kata

import (
	"regexp"
	"strings"
)

func NameValue(my_list []string) []int {
	result := make([]int, len(my_list))
	regex := regexp.MustCompile("[^a-zA-Z]")
	for i, v := range my_list {
		my_list[i] = strings.ToLower(v)
		my_list[i] = regex.ReplaceAllString(my_list[i], "")
		for j := 0; j < len(my_list[i]); j++ {
			result[i] += (int(my_list[i][j]) - 96) * (i + 1)
		}
	}
	return result
}
