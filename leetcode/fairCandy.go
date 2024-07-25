package leetcode

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum, bobSum := 0, 0
	aliceMap, bobMap := make(map[int]int), make(map[int]int)
	for _, v := range aliceSizes {
		aliceSum += v
		aliceMap[v]++
	}
	for _, v := range bobSizes {
		bobSum += v
		bobMap[v]++
	}

	for k, _ := range aliceMap {
		newSum := aliceSum - k
		newBob := bobSum + k

		diff := (newBob - newSum) / 2

		if _, ok := bobMap[diff]; ok {
			return []int{k, diff}
		}

	}
	return []int{}
}
