package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {

	sum := 0
	to := 0

	if start > destination {
		start, destination = destination, start
	}

	for k, v := range distance {
		if k < destination && k >= start {
			to += v
		}
		sum += v
	}

	return min(sum-to, to)

}
