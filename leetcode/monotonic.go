package leetcode

func isMonotonic(A []int) bool {
	inc := true
	dec := true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			dec = false
		}
		if A[i] < A[i-1] {
			inc = false
		}
		if !inc && !dec {
			return false
		}
	}
	return true
}
