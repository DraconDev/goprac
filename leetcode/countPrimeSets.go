package leetcode

func countPrimeSetBits(left int, right int) int {
	sets := 0
	for i := left; i <= right; i++ {
		bits := 0
		for j := 0; j < 32; j++ {
			if i&(1<<j) > 0 {
				bits++
			}
		}
		if bits == 2 || bits == 3 || bits == 5 || bits == 7 || bits == 11 || bits == 13 || bits == 17 || bits == 19 {
			sets++
		}
	}
	return sets
}
