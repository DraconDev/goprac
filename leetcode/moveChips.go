package leetcode

func minCostToMoveChips(position []int) int {
	odd := 0
	even := 0
	for _, v := range position {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd > even {
		return even
	}
	return odd

}
