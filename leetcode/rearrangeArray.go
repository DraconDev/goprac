package leetcode

func RearrangeArray(nums []int) []int {

	var pos, neg []int

	for _, num := range nums {
		if num < 0 {
			neg = append(neg, num)
		} else {
			pos = append(pos, num)
		}
	}
	var result []int

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			result = append(result, pos[i/2])
		} else {
			result = append(result, neg[i/2])
		}
	}
	return result

}
