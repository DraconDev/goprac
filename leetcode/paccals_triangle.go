package leetcode

func Pascals_triangle(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		temp := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				temp[j] = 1
			} else {
				temp[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = temp
	}
	return result
}
