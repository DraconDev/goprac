package dynamic

import "sort"

func MaxProfit(prices []int) int {
	profit := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit[i] {
				profit[i] = prices[j] - prices[i]
			}
		}
	}
	sort.Ints(profit)
	return profit[len(profit)-1]
}
