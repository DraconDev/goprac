package leetcode

func findDuplicates(nums []int) []int {
	var duplicates map[int]bool = make(map[int]bool)
	var result []int

	for _, v := range nums {
		if duplicates[v] {
			result = append(result, v)
		}
		duplicates[v] = true
	}
	return result
}
