package leetcode

func makeGood(s string) string {

	length := len(s)
	i := 0
	for i < length {
		if i+1 < length && abs(int(s[i])-int(s[i+1])) == 32 {
			s = s[:i] + s[i+2:]
			length -= 2
			i = 0
		} else {
			i++
		}
	}
	return s

}

func abs2(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
