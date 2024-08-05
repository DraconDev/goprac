package leetcode

func kthDistinct(arr []string, k int) string {
	countMap := make(map[string]int)
	order := []string{}
	for _, v := range arr {
		if _, ok := countMap[v]; !ok {
			order = append(order, v)
		}
		countMap[v]++
	}
	for v := range order {
		if countMap[order[v]] == 1 {
			k--
		}
		if k == 0 {
			return order[v]
		}
	}
	return ""

}
