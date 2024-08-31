package leetcode

func removeDuplicates(s string) string {

	runes := []rune(s)

	result := []rune{}
	for i := 0; i < len(runes); i++ {
		if len(runes) != i+1 && runes[i] == runes[i+1] {
			i++
		} else {
			if len(result) > 0 && result[len(result)-1] == runes[i] {
				result = result[:len(result)-1]
			} else {
				result = append(result, runes[i])
			}
		}
	}
	return string(result)

}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
