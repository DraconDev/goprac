package leetcode

func sortArrayByParity(nums []int) []int {
	odds := []int{}
	evens := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return append(evens, odds...)

}
