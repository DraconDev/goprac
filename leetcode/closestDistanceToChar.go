package leetcode

func ShortestToChar(s string, c byte) []int {
	letterLocation := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			letterLocation = append(letterLocation, i)
		}
	}

	result := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			result[i] = 0
		} else {
			minDistance := len(s)
			for _, j := range letterLocation {
				if j < i {
					minDistance = min(minDistance, i-j)
				} else {
					minDistance = min(minDistance, j-i)
				}
			}
			result[i] = minDistance
		}
	}

	return result
}
