package leetcode

func longestPalindrome(s string) int {

	letters := make(map[rune]int)
	for _, v := range s {
		letters[v]++
	}
	max := 0
	odd := 0
	for _, v := range letters {
		if v%2 == 1 {
			odd = 1
		}
		max += (v / 2) * 2
	}

	return max + odd
}
