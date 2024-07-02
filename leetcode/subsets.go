package leetcode

func Subsets(nums []int) [][]int {
	subsets := [][]int{{}}
	for _, num := range nums {
		for _, subset := range subsets {
			subsets = append(subsets, append([]int{}, append(subset, num)...))
		}
	}
	return subsets
}
