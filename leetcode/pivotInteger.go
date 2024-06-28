package leetcode

func PivotInteger(n int) int {
	var sum int
	if n%2 == 0 {
		sum = (1 + n) * n / 2

	} else {
		sum = ((1 + n) * n / 2)
	}

	var curr int
	for i := 0; i <= n; i++ {
		if curr == sum-curr-i {
			return i
		} else if curr > sum-curr-i {
			return -1
		}
		curr += i
	}

	return -1

}
