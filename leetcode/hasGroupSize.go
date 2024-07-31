package leetcode

func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	groupMap := make(map[int]int)
	for _, v := range deck {
		groupMap[v]++
	}

	count := 0
	for _, v := range groupMap {
		if v > count {
			count = v
		}
	}
outer:
	for i := 2; i <= count; i++ {
		for _, v := range groupMap {
			if v%i != 0 {
				continue outer
			}
		}
		return true
	}

	return false

}
