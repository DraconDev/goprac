package leetcode

func Anagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	word1 := make(map[rune]int)
	word2 := make(map[rune]int)
	for i, v := range s {
		word1[v]++
		word2[rune(t[i])]++
	}
	for k, v := range word1 {
		if word2[k] != v {
			return false
		}
	}
	return true
}
