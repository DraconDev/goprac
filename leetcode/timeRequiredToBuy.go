package leetcode

// func TimeRequiredToBuy(tickets []int, k int) int {

// 	if len(tickets) < k {
// 		return 0
// 	}
// 	var time int
// 	guest := &tickets[k]

// 	for *guest != 0 {
// 		if tickets[0] > 1 {
// 			tickets[0]--
// 			tickets = append(tickets, tickets[0])
// 			tickets = tickets[1:]
// 			time++
// 		} else {
// 			tickets = tickets[1:]
// 			time++
// 		}
// 	}
// 	return time
// }

func TimeRequiredToBuy(tickets []int, k int) int {

	if len(tickets) < k {
		return 0
	}
	var time int

	for tickets[k] != 0 {
		for i, v := range tickets {
			if v >= 1 {
				time++
			}
			tickets[i]--
			if i == k && tickets[k] == 0 {
				break
			}
		}
	}
	return time
}
