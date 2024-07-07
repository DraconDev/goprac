package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	count := 0
	remain := 0
	for numBottles > 0 {
		count += numBottles
		remain = numBottles % numExchange
		numBottles = numBottles/numExchange + remain
	}
	return count
}
