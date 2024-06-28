package leetcode

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	letters := make(map[byte]byte)
	letters2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := letters[s[i]]; ok {
			if letters[s[i]] != t[i] {
				return false
			}
		} else if _, ok := letters2[t[i]]; ok {
			if letters2[t[i]] != s[i] {
				return false
			}
		} else {
			letters[s[i]] = t[i]
			letters2[t[i]] = s[i]
		}
	}

	return true

}
