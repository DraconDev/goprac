package leetcode

func ReversePrefix(word string, ch byte) string {
	first := []byte{}
	second := []byte{}

	for i, v := range word {
		if string(v) == string(ch) {
			first = append(first, word[:i+1]...)
			second = append(second, word[i+1:]...)
			break
		}
	}
	// it empty
	if len(first) == 0 {
		return word
	}
	// reverse first
	for i := 0; i < len(first)/2; i++ {
		first[i], first[len(first)-i-1] = first[len(first)-i-1], first[i]
	}

	return string(first) + string(second)
}
