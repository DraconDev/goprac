package leetcode

func distributeCandies(candyType []int) int {
	// make a Set
	set := make(map[int]bool)
	// fill set
	for _, v := range candyType {
		set[v] = true
	}
	// get length
	return min(len(set), len(candyType)/2)
}
