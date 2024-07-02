package leetcode

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	var map1 = make(map[rune]int)
	var map2 = make(map[rune]int)

	for _, v := range word1 {
		map1[v]++
	}

	for _, v := range word2 {
		map2[v]++
	}

	if len(map1) != len(map2) {
		return false
	}

	var count = make(map[int]int)

	for k, v := range map1 {
		if map2[k] < 1 {
			return false
		}
		count[v]++
	}

	for k, v := range map2 {
		if map1[k] < 1 {
			return false
		}
		if count[v] < 1 {
			return false
		} else {
			count[v]--
		}

	}

	return true
}
