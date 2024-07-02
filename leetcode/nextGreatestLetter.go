package leetcode

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}
