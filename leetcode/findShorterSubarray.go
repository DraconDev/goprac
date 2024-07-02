package leetcode

func FindShortestSubArray(nums []int) int {

	numbersCount := map[int]int{}
	maxCount := 0

	for _, v := range nums {
		numbersCount[v]++
		if numbersCount[v] > maxCount {
			maxCount = numbersCount[v]
		}
	}
	visited := map[int]bool{}

	shortestSubarray := len(nums)

	firstIndex := 0
	lastIndex := len(nums)
	for _, letter := range nums {
		if visited[letter] {
			continue
		}
		for i, v := range nums {
			if numbersCount[letter] != maxCount {
				break
			}
			if v == letter {
				if !visited[letter] {
					visited[letter] = true
					firstIndex = i
					lastIndex = i
				} else {
					lastIndex = i
				}
			}
		}

		shortestSubarray = min(shortestSubarray, lastIndex-firstIndex+1)
		firstIndex = 0
		lastIndex = len(nums)
	}
	return shortestSubarray
}
