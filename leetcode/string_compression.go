package leetcode

import "strconv"

// func StringCompress(chars []byte) int {
// 	count := map[byte]int{}
// 	for _, v := range chars {
// 		count[v]++
// 	}
// 	result := []byte{}
// 	for k, v := range count {
// 		result = append(result, k)
// 		if v > 1 {
// 			result = append(result, strconv.Itoa(v)...)
// 		}
// 	}
// 	copy(chars, result)
// 	return len(result)
// }

func StringCompress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}

	currentLetter := chars[0]
	count := 1
	compressedChars := []byte{currentLetter}

	for i := 1; i < len(chars); i++ {
		if chars[i] != currentLetter {
			if count > 1 {
				countBytes := strconv.Itoa(count)
				compressedChars = append(compressedChars, []byte(countBytes)...)
			}

			currentLetter = chars[i]
			count = 1
			compressedChars = append(compressedChars, currentLetter)
		} else {
			count++
			if i == len(chars)-1 {
				if count > 1 {
					countBytes := strconv.Itoa(count)
					compressedChars = append(compressedChars, []byte(countBytes)...)
				}
			}
		}
	}

	chars = append(chars[:0], compressedChars...)
	return len(compressedChars)
}
