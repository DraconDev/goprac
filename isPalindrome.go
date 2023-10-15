package main

import "strconv"

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
