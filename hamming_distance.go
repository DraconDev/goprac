package main

// func HammingDistance(x int, y int) int {
// 	convertedX := strconv.FormatInt(int64(x), 2)
// 	convertedY := strconv.FormatInt(int64(y), 2)
// 	if len(convertedX) > len(convertedY) {
// 		convertedX, convertedY = convertedY, convertedX
// 	}
// 	difference := 0
// 	gap := len(convertedY) - len(convertedX)
// 	for i, v := range convertedY {
// 		if i < gap {
// 			if v == '1' {
// 				difference++
// 			}
// 		} else if v != rune(convertedX[i-gap]) {
// 			difference++
// 		}

// 	}
// 	return difference
// }

// func HammingDistance(x int, y int) int {
// 	short := strconv.FormatInt(int64(x), 2)
// 	long := strconv.FormatInt(int64(y), 2)
// 	if len(short) > len(long) {
// 		short, long = long, short
// 	}
// 	difference := 0
// 	for i := 0; i < len(long)-len(short); i++ {
// 		if long[i] == '1' {
// 			difference++
// 		}
// 	}
// 	long = long[len(long)-len(short):]
// 	for i, v := range long {
// 		if v != rune(short[i]) {
// 			difference++
// 		}

// 	}
// 	return difference
// }

func HammingDistance(x, y int) int {
	var distance int
	for x != 0 || y != 0 {
		x &= x - 1
		y &= y - 1
		distance++
	}
	return distance
}
