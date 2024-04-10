package leetcode

import "sort"

func DeckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	Queue := []int{}
	for i := len(deck) - 1; i >= 1; i-- {
		Queue = append(Queue, deck[i])
		Queue = append(Queue, Queue[0])
		Queue = Queue[1:]
	}
	Queue = append(Queue, deck[0])
	for i, j := 0, len(Queue)-1; i < j; i, j = i+1, j-1 {
		Queue[i], Queue[j] = Queue[j], Queue[i]
	}
	return Queue
}
