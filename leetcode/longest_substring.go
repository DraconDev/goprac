package leetcode

func LengthOfLongestSubstring(s string) int {
	var window = []byte{}

	max := 0

	for e := range s {
		for i, v := range window {
			if v == s[e] {
				if max < len(window) {
					max = len(window)
				}
				window = window[i+1:]
				break
			}
		}
		window = append(window, s[e])
	}

	if max < len(window) {
		max = len(window)
	}

	return max

}
