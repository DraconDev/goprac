package leetcode

func reverseString(s []byte) {
	for i, _ := range s[:len(s)/2] {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
