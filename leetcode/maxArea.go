package leetcode

func MaxArea(height []int) int {
	var l, r = 0, len(height) - 1
	var maxArea = 0
	for l < r {
		tempArea := min(height[l], height[r]) * (len(height) - l - (len(height) - r))
		if tempArea > maxArea {
			maxArea = tempArea
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
