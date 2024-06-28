package leetcode

func SelfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividing(num int) bool {
	for num > 0 {
		if num%10 == 0 || num%(num%10) != 0 {
			return false
		}
		num = num / 10
	}
	return true
}
