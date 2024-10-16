package leetcode

func romanToInt(s string) int {

	roman_map := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman_map[string(s[i])] < roman_map[string(s[i+1])] {
			result -= roman_map[string(s[i])]
		} else {
			result += roman_map[string(s[i])]
		}
	}

	return result
}
