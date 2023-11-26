package leetcode

func LargestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := []int{-1}
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]

			stack = stack[:len(stack)-1]

			width := i - stack[len(stack)-1] - 1

			area := height * width
			if area > maxArea {
				maxArea = area
			}

		}
		stack = append(stack, i)
	}
	return maxArea
}
