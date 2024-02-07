package leetcode

func FrequencySort(s string) string {

	length := 0
	var freq = make(map[rune]int)
	for _, v := range s {
		freq[v]++
		if freq[v] > length {
			length = freq[v]
		}
	}

	reverse := make([][]rune, length+1)
	for k, v := range freq {
		reverse[v] = append(reverse[v], k)
	}

	result := ""
	for i := length; i > 0; i-- {
		for _, v := range reverse[i] {
			for j := 0; j < i; j++ {
				result += string(v)
			}
		}
	}

	return result
}
