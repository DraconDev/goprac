package kata

import "unicode"

func ReverseLetters(s string) string {
	arr := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			arr = append([]rune{v}, arr...)
		}
	}
	return string(arr)

}
