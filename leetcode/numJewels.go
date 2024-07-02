package leetcode

func numJewelsInStones(jewels string, stones string) int {
	jewel := map[byte]bool{}
	count := 0

	for i := 0; i < len(jewels); i++ {
		jewel[jewels[i]] = true
	}
	for i := 0; i < len(stones); i++ {
		if jewel[stones[i]] {
			count++
		}
	}
	return count

}
