package main

import "strings"

func Capitalize(st string, arr []int) string {
	//..
	slice := []byte(st)
	for _, v := range arr {
		// slice[v] = strings.ToUpper(string(slice[v]))
		length := len(slice)
		if v < length {
			slice[v] = strings.ToUpper(string(slice[v]))[0]
		}
	}
	return string(slice)
}
