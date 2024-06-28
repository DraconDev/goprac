package leetcode

// func FindLUSlength(a string, b string) int {

// 	if a == b {
// 		return -1
// 	}

// 	var max int
// 	var currentMax int
// 	for i := 0; i < len(a); i++ {
// 		if a[i] != b[i] {
// 			currentMax++
// 		} else {
// 			if currentMax > max {
// 				max = currentMax
// 			}
// 			currentMax = 0
// 		}
// 	}
// 	if currentMax > max {
// 		max = currentMax
// 	}

// 	return max
// }

func FindLUSlength(a string, b string) int {

	if a == b {
		return -1
	}

	max := Max(len(a), len(b))

	return max
}
