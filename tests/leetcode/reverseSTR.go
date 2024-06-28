package leetcode

func ReverseStr(s string, k int) string {

	firstByte := []byte(s)

	result := []byte{}
	// reverse first half

	first := true
	for i := 0; i < len(firstByte); i += k {

		if first {
			first = false
			for j := min(i+k-1, len(firstByte)-1); j >= i; j-- {
				result = append(result, firstByte[j])
			}
		} else {
			first = true
			for j := i; j < min(i+k, len(firstByte)); j++ {
				result = append(result, firstByte[j])
			}
		}

	}

	return string(result)

}
