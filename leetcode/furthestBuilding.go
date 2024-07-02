package leetcode

import "container/heap"

// func FurthestBuilding(heights []int, bricks int, ladders int) int {
// 	var canReach int
// 	for i, v := range heights[1:] {
// 		if v <= heights[i] {
// 			canReach++
// 			continue
// 		}
// 		if ladders > 0 && bricks == v-heights[i] {
// 			ladders--
// 		} else if v-heights[i] <= bricks {
// 			bricks -= v - heights[i]
// 		} else if ladders > 0 {
// 			ladders--
// 		} else {
// 			break
// 		}
// 		canReach++

// 	}
// 	return canReach

// }

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Top() int {
	return (*h)[0]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	pq := make(IntHeap, 0, ladders)
	heap.Init(&pq)
	n := len(heights)
	for i := 0; i < n; i++ {
		if i == n-1 {
			return i
		}
		if heights[i] >= heights[i+1] {
			continue
		}

		brickNeeds := heights[i+1] - heights[i]
		if pq.Len() < ladders {
			heap.Push(&pq, brickNeeds)
		} else if pq.Len() > 0 {
			// ladders used up
			if brickNeeds > pq.Top() {
				bricks -= pq.Top()
				heap.Pop(&pq)
				heap.Push(&pq, brickNeeds)
			} else {
				bricks -= brickNeeds
			}
		} else {
			bricks -= brickNeeds
		}

		if bricks < 0 {
			return i
		}
	}

	return n - 1
}

// if ladders > 0 && bricks == v-heights[i] {
// 	ladders--
// } else if v-heights[i] <= bricks {
// 	bricks -= v - heights[i]
// } else if ladders > 0 {
// 	ladders--
// } else {
// 	break
// }
