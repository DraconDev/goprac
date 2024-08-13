package leetcode

func isAlienSorted(words []string, order string) bool {
	orderIndex := 0

	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	for i := 0; i < len(words); i++ {
		if orderIndex >= len(order) {
			return false
		}
		for words[i][0] != order[orderIndex] {
			if orderIndex == len(order)-1 {
				return false
			}
			orderIndex++
		}
		wordsWithSameLetter := []string{words[i]}
		for j := i + 1; j < len(words); j++ {
			if words[j][0] == order[orderIndex] {
				wordsWithSameLetter = append(wordsWithSameLetter, words[j])
			}
		}

		if !checkIfWordsWithSameLetterAreSorted(wordsWithSameLetter, orderMap) {
			return false
		}
		orderIndex++
		i += len(wordsWithSameLetter) - 1

	}

	return true
}

func checkIfWordsWithSameLetterAreSorted(wordsWithSameLetter []string, orderMap map[byte]int) bool {

	for i := 0; i < len(wordsWithSameLetter)-1; i++ {
		for j := 0; j < len(wordsWithSameLetter[i]); j++ {
			if j >= len(wordsWithSameLetter[i+1]) {
				return false
			}
			if orderMap[wordsWithSameLetter[i][j]] < orderMap[wordsWithSameLetter[i+1][j]] {
				return true
			}
			if orderMap[wordsWithSameLetter[i][j]] > orderMap[wordsWithSameLetter[i+1][j]] {
				return false
			}
		}
	}
	return true
}
