package leetcode

func CommonChars(words []string) []string {
	// array of maps
	letters := make([]map[rune]int, len(words))
	for i, word := range words {
		letters[i] = make(map[rune]int)
		for _, c := range word {
			letters[i][c]++
		}
	}

	common := []string{}

	for k, v := range letters[0] {
		for j, word2 := range letters[1:] {
			if letters[j+1] == nil {
				v = 0
				break
			} else {
				if v > word2[k] {
					v = word2[k]
				}
			}
		}
		for i := 0; i < v; i++ {
			common = append(common, string(k))
		}
	}
	return common
}
