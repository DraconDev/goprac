package leetcode

func findDisappearedNumbers(nums []int) []int {

	numbers := make(map[int]bool)
	length := len(nums)

	if length == 0 {
		return []int{}
	}

	for _, v := range nums {
		numbers[v] = true
	}

	missing := []int{}
	for i := 1; i <= length; i++ {
		if numbers[i] == true {
			continue
		}
		missing = append(missing, i)
	}

	return missing

}
