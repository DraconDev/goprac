package leetcode

import "sort"

func minimumPushes(word string) int {
	lettersMap := make(map[rune]int)

	for _, v := range word {
		lettersMap[v]++
	}

	presses := []int{}

	for _, v := range lettersMap {
		presses = append(presses, v)
	}
	sort.Slice(presses, func(i, j int) bool {
		return presses[i] > presses[j]
	})

	result := 0

	for i := 0; i < len(presses); i++ {
		if i >= 24 {
			result += presses[i] * 4
		} else if i >= 16 {
			result += presses[i] * 3
		} else if i >= 8 {
			result += presses[i] * 2
		} else {
			result += presses[i]
		}
	}

	return result
}
