package main

func romanToInt(s string) int {
	result := 0
	roman_map := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := 0; i < len(s)-1; i++ {
		if roman_map[string(s[i])] < roman_map[string(s[i+1])] {
			result -= roman_map[string(s[i])]
		} else {
			result += roman_map[string(s[i])]
		}
	}
	result += roman_map[string(s[len(s)-1])]
	return result

}
