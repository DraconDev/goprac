package main

import (
	"strings"
)

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

func isInSlice(slice []int, val int) bool {
	for _, v := range slice {
		check := v * v
		// check := int(math.Sqrt(float64(val)))
		if check == val {
			return true
		}
	}
	return false
}

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil && array2 == nil {
		return false
	}
	for _, v := range array2 {
		if !isInSlice(array1, v) {
			return false
		}
	}
	return true
}
