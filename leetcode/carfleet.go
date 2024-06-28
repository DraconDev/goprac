package leetcode

import (
	"sort"
)

// func CarFleet(target int, position []int, speed []int) int {
// 	timeArr := map[int]float64{}

// 	for i := 0; i < len(position); i++ {
// 		timeArr[position[i]] = float64(target-position[i]) / float64(speed[i])
// 	}

// 	sort.Slice(position, func(i, j int) bool {
// 		return position[i] < position[j]
// 	})

// 	for i := 0; i < len(position)-1; i++ {
// 		for j := i + 1; j <= len(position)-1; j++ {
// 			if timeArr[position[i]] <= timeArr[position[j]] {
// 				delete(timeArr, position[i])
// 				break
// 			}
// 		}
// 	}

// 	return len(timeArr)
// }

func CarFleet(target int, position []int, speed []int) int {
	timeMap := make(map[int]float64)

	for i := 0; i < len(position); i++ {
		timeMap[-position[i]] = float64(target-position[i]) / float64(speed[i])
	}

	fleetCount := 0
	maxTime := 0.0

	keys := make([]int, len(timeMap))
	i := 0
	for k := range timeMap {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	for _, k := range keys {
		if timeMap[k] > maxTime {
			maxTime = timeMap[k]
			fleetCount++
		}
	}

	return fleetCount
}
