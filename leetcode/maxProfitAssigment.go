package leetcode

import "sort"

type Job struct {
	Profit     int
	Difficulty int
}

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	result := 0
	sort.Ints(worker)

	// zip profit and difficulty
	zipped := make([]struct{ Profit, Difficulty int }, len(difficulty))

	for i := 0; i < len(difficulty); i++ {
		zipped[i] = Job{profit[i], difficulty[i]}
	}

	sort.Slice(zipped, func(i, j int) bool {
		return zipped[i].Profit > zipped[j].Profit
	})

	for _, w := range zipped {
		if len(worker) == 0 {
			break
		}
		for i := len(worker) - 1; i >= 0; i-- {
			if w.Difficulty <= worker[i] {
				result += w.Profit
				worker = worker[:i]
			} else {
				break
			}
		}
	}

	return result

}
