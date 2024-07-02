package leetcode

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n%4 != 0 || n < 1 {
		return false
	}
	return isPowerOfFour(n / 4)
}
