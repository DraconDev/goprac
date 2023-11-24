package leetcode

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	timeArr := map[int]float64{}

	for i := 0; i < len(position); i++ {
		timeArr[position[i]] = float64(target-position[i]) / float64(speed[i])
	}

	sort.Slice(position, func(i, j int) bool {
		return position[i] < position[j]
	})

	for i := 0; i < len(position)-1; i++ {
		for j := i + 1; j <= len(position)-1; j++ {
			if timeArr[position[i]] <= timeArr[position[j]] {
				delete(timeArr, position[i])
				break
			}
		}
	}

	return len(timeArr)
}

// func CarFleet(target int, position []int, speed []int) int {
// 	timeArr := make([]float64, len(position))

// 	for i := 0; i < len(position); i++ {
// 		timeArr[i] = float64(target-position[i]) / float64(speed[i])
// 	}

// 	sort.Slice(position, func(i, j int) bool {
// 		return position[i] < position[j]
// 	})

// 	fleetCount := 0
// 	maxTime := 0.0

// 	for i := 0; i < len(position); i++ {
// 		if timeArr[i] > maxTime {
// 			fleetCount++
// 			maxTime = timeArr[i]
// 		} else {
// 			timeArr[i] = maxTime
// 		}
// 	}

// 	return fleetCount
// }
