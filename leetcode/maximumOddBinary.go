package leetcode

func maximumOddBinaryNumber(s string) string {

	var count int

	for _, v := range s {
		if v == '1' {
			count++
		}

	}
	result := ""

	for i := 0; i < count-1; i++ {
		result += "1"
	}
	for i := 0; i < len(s)-count; i++ {
		result += "0"
	}
	result += "1"
	return result

}
