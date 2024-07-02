package main

func Grow(arr []int) int {
	sum := 1
	for _, v := range arr {
		sum *= v
	}
	return sum
}
