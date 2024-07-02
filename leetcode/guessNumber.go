package leetcode

func guessNumber(n int) int {
	start, end := 1, n
	for start <= end {
		mid := (start + end) / 2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		default:
			return mid
		}
	}
	return 0
}

func guess(num int) int {
	if num == 6 {
		return 0
	} else if num < 6 {
		return -1
	} else {
		return 1
	}
}
