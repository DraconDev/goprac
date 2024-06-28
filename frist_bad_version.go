package main

func isBadVersion(version int) bool {
	// test := 4
	return version == 4
}

// func firstBadVersion(n int) int {
// 	for n > 1 {
// 		if isBadVersion(n) {
// 			return n
// 		}
// 		n--
// 	}
// 	return 1
// }

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
