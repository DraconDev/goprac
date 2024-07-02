package leetcode

// func tribonacci(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	if n == 2 {
// 		return 1
// 	}
// 	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
// }

func Tribonacci(n int) int {
	var sum []int

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			sum = append(sum, i)
		} else if i == 2 {
			sum = append(sum, 1)
		} else {
			sum = append(sum, sum[i-1]+sum[i-2]+sum[i-3])
		}
	}
	return sum[n]

}
