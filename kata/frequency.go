package kata

import (
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) string {
	arr := []string{}
	for _, v := range str {
		arr = append(arr, strconv.Itoa(strings.Count(str, string(v))))
	}
	result := strings.Join(arr, sep)
	return result //Code goes here!
}
