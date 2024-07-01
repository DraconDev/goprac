package leetcode

func threeConsecutiveOdds(arr []int) bool {
	odd := 0
	for _, v := range arr {
		if v%2 != 0 {
			odd++
			if odd == 3 {
				return true
			}
		} else {
			odd = 0
		}
	}
	return false
}
