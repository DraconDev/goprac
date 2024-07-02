package leetcode

import (
	"sort"
	"strconv"
)

func FindRelativeRanks(score []int) []string {

	// copy of score
	ranking := make([]int, len(score))
	// fill
	copy(ranking, score)

	// sort descending
	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i] > ranking[j]
	})

	// assing rank to each score
	ranks := make(map[int]string)

	for i, v := range ranking {
		if i == 0 {
			ranks[v] = "Gold Medal"
		} else if i == 1 {
			ranks[v] = "Silver Medal"
		} else if i == 2 {
			ranks[v] = "Bronze Medal"
		} else {
			ranks[v] = strconv.Itoa(i + 1)
		}
	}

	// return ranks
	var result []string

	for _, v := range score {
		result = append(result, ranks[v])
	}

	return result

}
