package leetcode

func subtractProductAndSum(n int) int {
	res := 1
	sum := 0

	for n > 0 {
		digit := n % 10
		res *= digit
		sum += digit
		n /= 10
	}

	return res - sum

}
