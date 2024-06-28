package leetcode

func MaxVowels(s string, k int) int {
	currentMax, max := 0, 0

	// * count vowels
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currentMax++
		}
	}
	// * update max
	if currentMax == k {
		return currentMax
	}
	max = currentMax

	// * find true max
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			currentMax++
		}
		if isVowel(s[i-k]) {
			currentMax--
		}
		if currentMax == k {
			return currentMax
		}
		if currentMax > max {
			max = currentMax
		}
	}
	return max
}

func isVowel(char byte) bool {
	// var vowelLookup = map[byte]bool{
	// 	'a': true,
	// 	'e': true,
	// 	'i': true,
	// 	'o': true,
	// 	'u': true,
	// 	'A': true,
	// 	'E': true,
	// 	'I': true,
	// 	'O': true,
	// 	'U': true,
	// }
	test := rune(char)
	switch test {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
