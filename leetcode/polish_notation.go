package leetcode

import "strconv"

func EvalRPN(tokens []string) int {
	stack := []int{}

	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			calculateOperation(v, &stack)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func calculateOperation(s string, stack *[]int) {
	// * if operation
	if s == "+" {
		(*stack)[len(*stack)-2] += (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
	} else if s == "-" {
		(*stack)[len(*stack)-2] -= (*stack)[len(*stack)-1] // Subtract instead of returning
		*stack = (*stack)[:len(*stack)-1]
	} else if s == "*" {
		(*stack)[len(*stack)-2] *= (*stack)[len(*stack)-1] // Multiply instead of returning
		*stack = (*stack)[:len(*stack)-1]
	} else if s == "/" {
		(*stack)[len(*stack)-2] /= (*stack)[len(*stack)-1] // Divide instead of returning
		*stack = (*stack)[:len(*stack)-1]
	}
}
