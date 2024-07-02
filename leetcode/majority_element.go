package leetcode

func majorityElement(nums []int) int {
	var elements = make(map[int]int)

	for _, v := range nums {
		elements[v]++
		if elements[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
