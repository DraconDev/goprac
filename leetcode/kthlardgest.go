package leetcode

import "sort"

type KthLargest struct {
	k       int
	numbers []int
}

func Constructor1(k int, nums []int) KthLargest {
	return KthLargest{k: k, numbers: nums}
}

func (this *KthLargest) Add(val int) int {

	this.numbers = append(this.numbers, val)
	sort.Ints(this.numbers)
	return this.numbers[len(this.numbers)-this.k]
}
