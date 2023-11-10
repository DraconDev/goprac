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

// func LettersMap(str string) map[rune]int {
// 	result := make(map[rune]int)
// 	for _, v := range str {
// 		result[v]++
// 	}
// 	return result
// }

// func GroupAnagrams(strs []string) [][]string {
// 	result := [][]string{}
// 	var anagrams [][map[rune]int,string]
// 	for _, v := range strs {
// 		anagrams = append(anagrams, LettersMap(v))
// 	}
// 	for _, v := range anagrams {
// 		for k, _ := range v {
// 			temp = append(temp, string(k))
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }

// func GroupAnagrams(strs []string) [][]string {
// 	// sort array elements
// 	var anagrams []string

// 	for _, v := range strs {
// 		anagrams = append(anagrams, sort.Strings(v))
// 	}

// 	// rest of the code
// 	// ...

// 	// make a Set
// 	// add to array if anagram to set key
// 	// sort array from smallest to biggest
// 	return [][]string{}

// }
