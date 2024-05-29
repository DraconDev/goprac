package leetcode

import "strconv"

func calPoints(operations []string) int {
	stack := []int{}
	for _, v := range operations {
		if v == "C" {
			stack = stack[:len(stack)-1]
		} else if v == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		} else if v == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}
