package leetcode

func arrangeCoins(n int) int {

	i := 1

	for n-i >= 0 {
		n -= i
		i++
	}
	return i - 1

}
