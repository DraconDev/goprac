package leetcode

import "strings"

// func licenseKeyFormatting(s string, k int) string {
// 	s = strings.ToUpper(s)

// 	prefix := ""
// 	// cut out first dash
// 	for i, v := range s {
// 		if v == '-' {
// 			prefix = s[:i]
// 			s = s[i+1:]
// 			break
// 		}
// 	}

// 	// filter out dash
// 	s = strings.Replace(s, "-", "", -1)

// 	result := prefix
// 	// add rest of the letters
// 	for i := 0; i < len(s); i += k {
// 		if i+k > len(s) {
// 			result += "-" + s[i:]
// 		} else {
// 			result += "-" + s[i:i+k]
// 		}
// 	}

// 	return result

// }

func LicenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(s)

	// filter out dash
	s = strings.Replace(s, "-", "", -1)

	result := ""
	// add rest of the letters
	for i := len(s); i > 0; i -= k {
		if i > k {
			result = "-" + s[i-k:i] + result
		} else {
			result = s[0:i] + result
		}
	}

	return result
}
