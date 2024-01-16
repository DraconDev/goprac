package leetcode

import "math/rand"

type RandomizedSet struct {
	set   map[int]int // key: value, value: index
	items []int       // value
}

func Const() RandomizedSet {
	return RandomizedSet{
		set:   make(map[int]int),
		items: make([]int, 0),
	}

}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.items = append(this.items, val)
	this.set[val] = len(this.items) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {

	if _, ok := this.set[val]; !ok {
		return false
	}
	if this.set[val] == len(this.items)-1 {
		this.items = this.items[:len(this.items)-1]
	} else {
		// update
		this.set[this.items[len(this.items)-1]] = this.set[val]
		// swap
		this.items[this.set[val]] = this.items[len(this.items)-1]
		// delete
		this.items = this.items[:len(this.items)-1]

	}

	delete(this.set, val)
	return true

}

func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.items))

	return this.items[random]
}
