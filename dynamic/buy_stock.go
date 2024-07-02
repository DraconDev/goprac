package dynamic

// func MaxProfit(prices []int) int {
// 	profit := make([]int, len(prices))
// 	for i := 0; i < len(prices); i++ {
// 		for j := i + 1; j < len(prices); j++ {
// 			if prices[j]-prices[i] > profit[i] {
// 				profit[i] = prices[j] - prices[i]
// 			}
// 		}
// 	}
// 	sort.Ints(profit)
// 	return profit[len(profit)-1]
// }

func MaxProfit(prices []int) int {
	var l, r = 0, 1
	profit := 0
	for r < len(prices) {
		if prices[r]-prices[l] > profit {
			profit = prices[r] - prices[l]
		}
		if prices[r]-prices[l] < 0 {
			l = r
		}
		r++
	}
	return profit
}
