package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := findMax(candies)
	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if (candies[i] + extraCandies) >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
