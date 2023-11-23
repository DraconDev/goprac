package leetcode

// func dailyTemperatures(temperatures []int) []int {
// 	days := make([]int, len(temperatures))
// 	for i, v := range temperatures {
// 		for j := i + 1; j < len(temperatures); j++ {
// 			if temperatures[j] > v {
// 				days[i] = j - i
// 				break
// 			}
// 		}
// 	}
// 	return days
// }

// func dailyTemperatures(temperatures []int) []int {
// 	days := make([]int, len(temperatures))
// 	for i, v := range temperatures {
// 		for j := i + 1; j < len(temperatures); j++ {
// 			if temperatures[j] > v {
// 				days[i] = j - i
// 				break
// 			}
// 		}
// 	}
// 	return days
// }

func DailyTemperatures(temperatures []int) []int {
    // Initialize an empty stack to store indices of temperatures
    stack := []int{}
    // Initialize an array to store the result
    result := make([]int, len(temperatures))

    for i := 0; i < len(temperatures); i++ {
        // Check if the current temperature is greater than the temperature at the top of the stack
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            // If it is, calculate the number of days until a warmer temperature is found
            prevIndex := stack[len(stack)-1]
            result[prevIndex] = i - prevIndex
            // Pop the top element from the stack
            stack = stack[:len(stack)-1]
        }
        // Push the current index onto the stack
        stack = append(stack, i)
    }

    return result
}