package leetcode

import "fmt"

// func GCDOfStrings(str1 string, str2 string) string {
// 	var shorter, longer = str1, str2
// 	if len(str1) > len(str2) {
// 		shorter = str2
// 		longer = str1
// 	}
// 	longest := ""
// 	for i := 0; i < len(shorter) && len(longest) < len(shorter)-i; i++ {
// 		temp := ""
// 		for j := 0; j < len(shorter)-i; j++ {
// 			if longer[j] != shorter[j] {
// 				break
// 			}
// 			temp += string(longer[j])
// 		}
// 		if len(temp) > len(longest) {
// 			longest = temp
// 		}
// 	}
// 	return longest
// }

// func GCDOfStrings(str1 string, str2 string) string {
// 	if str1+str2 != str2+str1 {
// 		return ""
// 	}
// 	if str1 == str2 {
// 		return str1
// 	}
// 	var shorter, longer = str1, str2
// 	if len(str1) > len(str2) {
// 		shorter = str2
// 		longer = str1
// 	}
// 	longest := ""
// 	for i := len(shorter); i > 0; i-- {
// 		temp := shorter[0:i]
// 		// for j := 0; j < len(shorter)-i; j++ {

// 		// }
// 		long_length := len(longer)
// 		for long_length > 0 {
// 			if longer[0:i] == shorter {

// 			}
// 		}

// 	}
// 	return longest
// }

func GCDOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	if str1 == str2 {
		return str1
	}
	var shorter, longer = str1, str2
	if len(str1) > len(str2) {
		shorter = str2
		longer = str1
	}
	for i := len(shorter); i > 0; i-- {
		// * make divisor slice
		divisor := shorter[0:i]
		// * check if divides long
		for j := 0; j < len(longer); j += len(divisor) {
			if len(shorter)%len(divisor) != 0 {
				divisor = ""
				break
			}
			if j+i > len(longer) {
				divisor = ""
				break
			}
			if longer[j:j+i] != divisor {
				divisor = ""
				break
			}
			if j+i < len(shorter) {
				if shorter[j:j+i] != divisor {
					divisor = ""
					break
				}
			}
		}
		// * if not empty then best result
		if divisor != "" {
			fmt.Printf("divisor: %s\n", divisor)
			return divisor
		}
	}
	return ""
}
