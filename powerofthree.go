package main

// func isPowerOfThree(n int) bool {
// 	if n == 3 {
// 		return true
// 	}
// 	if n%3 == 0 {
// 		return isPowerOfThree(n / 3)
// 	}
// 	return false
// }

func isPowerOfThree(n int) bool {
	for i := 3; i < n; i *= 3 {
		if i == n {
			return true
		}
	}
	return false
}
