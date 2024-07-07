package leetcode

func NumWaterBottles(numBottles int, numExchange int) int {
	count := 0
	empty := 0
	for numBottles > 0 {
		count += numBottles
		empty += numBottles
		refill := empty / numExchange
		empty -= refill * numExchange
		numBottles = refill
	}
	return count
}
