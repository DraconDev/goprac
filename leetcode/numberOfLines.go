package leetcode

func NumberOfLines(widths []int, s string) []int {
	pixelsSum := 0
	rows := 1
	for _, v := range s {
		if pixelsSum+widths[v-'a'] > 100 {
			rows++
			pixelsSum = widths[v-'a']
		} else {
			pixelsSum += widths[v-'a']
		}
	}

	return []int{rows, pixelsSum}
}
