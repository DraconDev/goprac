package leetcode

func CountBinarySubstrings(s string) int {
	subStrings := make(map[string]int)

	for i := range s {
		count := 0

		flip := struct {
			c rune
			f bool
		}{'0', false}

		for j, c := range s[i:] {

			if c == '0' {
				if j == 0 {
					flip.c = '0'
					flip.f = false
				}
				if flip.c == '1' {
					flip.f = true
				}
				if flip.c == '0' && flip.f {
					break
				}

				count--
			} else {
				if j == 0 {
					flip.c = '1'
					flip.f = false
				}
				if flip.c == '0' {
					flip.f = true
				}
				if flip.c == '1' && flip.f {
					break
				}
				count++
			}
			if count == 0 {
				subStrings[s[i:i+j+1]]++
				break
			}
		}

	}
	result := 0

	for _, v := range subStrings {
		result += v
	}
	return result

}
