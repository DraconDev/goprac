package leetcode

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	difference := 0
	letters := make(map[byte]int)
	duplicate := false
	differentLettersA := []byte{}
	differentLettersB := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			difference++
			if difference > 2 {
				return false
			}
			differentLettersA = append(differentLettersA, s[i])
			differentLettersB = append(differentLettersB, goal[i])
		}
		if !duplicate {
			letters[s[i]]++
			if letters[s[i]] > 1 {
				duplicate = true
			}
		}

	}
	if difference == 1 {
		return false
	}
	if difference == 0 && duplicate {
		return true
	}
	if difference == 2 {
		return differentLettersA[0] == differentLettersB[1] && differentLettersA[1] == differentLettersB[0]
	}
	return false

}
