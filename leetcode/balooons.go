package leetcode

func maxNumberOfBalloons(text string) int {

	type letter struct {
		letter byte
		count  int
	}

	findLetters := []letter{
		{'b', 1},
		{'a', 1},
		{'l', 2},
		{'o', 2},
		{'n', 1},
	}

	letters := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		letters[text[i]]++
	}

	result := len(text)

	for i := 0; i < len(findLetters); i++ {
		if letters[findLetters[i].letter] > 0 {
			result = min(letters[findLetters[i].letter]/findLetters[i].count, result)
		} else {
			return 0
		}
	}
	return result

}
