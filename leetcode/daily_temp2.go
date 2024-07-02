package leetcode

func DailyTemperatures2(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int

	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

	}

	return result
}
