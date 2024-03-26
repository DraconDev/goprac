package leetcode

func findDisappearedNumbers(nums []int) []int {

	numbers := make(map[int]bool)
	length := len(nums)

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
