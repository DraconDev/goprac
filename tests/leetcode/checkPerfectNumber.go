package leetcode

func checkPerfectNumber(num int) bool {

	var sum int

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num

}
