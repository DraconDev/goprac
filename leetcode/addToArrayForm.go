package leetcode

func addToArrayForm(num []int, k int) []int {
	carry := k
	result := []int{}
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + carry
		result = append(result, sum%10)
		carry = sum / 10
	}
	for carry > 0 {
		result = append(result, carry%10)
		carry /= 10
	}
	// reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
