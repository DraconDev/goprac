package leetcode

// func SumSubarrayMins(arr []int) int {
// 	var result int

// 	for i := 0; i < len(arr); i++ {
// 		min := math.MaxInt
// 		for j := i; j < len(arr); j++ {
// 			if arr[j] < min {
// 				min = arr[j]
// 			}
// 			result += min
// 		}
// 	}

// 	return result % (1000000000 + 7)

// }

func SumSubarrayMins(arr []int) int {
	mod := int(1e9) + 7 // Define the modulo value for the result.
	var stack []int     // Stack to store indices of elements.
	var result int      // Accumulator for the sum of minimums.

	// Iterate over all elements of the array, and one additional iteration for the 'virtual' right end.
	for right := 0; right <= len(arr); right++ {
		// Process the elements on the stack. The condition checks if we're at the 'virtual' right end
		// or the current element is less than or equal to the element at the top of the stack.
		for len(stack) > 0 && (right == len(arr) || arr[stack[len(stack)-1]] > arr[right]) {
			// Pop the top element index from the stack.
			midIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Determine the left boundary index for the current subarray.
			leftIndex := -1
			if len(stack) > 0 {
				leftIndex = stack[len(stack)-1]
			}

			// Calculate the number of subarrays where arr[midIndex] is the minimum.
			widthRight := right - midIndex    // Number of possibilities to the right of midIndex.
			widthLeft := midIndex - leftIndex // Number of possibilities to the left of midIndex.

			// Update the result with the contribution of subarray minimums for arr[midIndex].
			result += (widthRight * widthLeft * arr[midIndex]) % mod
			result %= mod // Ensure result stays within bounds of mod.
		}

		// Push the current index onto the stack.
		stack = append(stack, right)
	}

	return result // Return the accumulated result.
}
