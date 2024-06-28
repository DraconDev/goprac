package main

// func ReverseBits(num uint32) uint32 {
// 	converted := strconv.FormatUint(uint64(num), 2)
// 	runed := []rune(converted)
// 	for i := 0; i < len(runed)/2; i++ {
// 		runed[i], runed[len(runed)-i-1] = runed[len(runed)-i-1], runed[i]
// 	}
// 	result, _ := strconv.ParseUint(string(runed), 2, 32)

// 	return uint32(result)
// }

// func ReverseBits(num uint32) uint32 {
// 	converted := strconv.FormatUint(uint64(num), 2)
// 	runed := []rune(converted)
// 	result := []rune{}
// 	for i := len(runed) - 1; i >= 0; i-- {
// 		result = append(result, runed[i])
// 	}
// 	reversedNum, _ := strconv.ParseUint(string(result), 2, 32)
// 	return uint32(reversedNum)
// }

func ReverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num = num >> 1
	}
	return result
}

// func ReverseBits(num uint32) uint32 {
// 	var result uint32
// 	for i := 0; i < 32; i++ {
// 		// Shift the bits of 'result' one position to the left
// 		result = (result << 1)

// 		// Get the least significant bit of 'num' using bitwise AND
// 		leastSignificantBit := num & 1

// 		// Set the least significant bit of 'result' to the value of the least significant bit of 'num'
// 		result = result | leastSignificantBit

// 		// Shift the bits of 'num' one position to the right
// 		num = num >> 1
// 	}
// 	return result
// }
