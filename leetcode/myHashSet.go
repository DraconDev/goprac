package leetcode

type MyHashSet struct {
	set [125001]byte
}

func Constructor3() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	// find the byte to store the key
	byt := key / 8
	bit := key % 8
	this.set[byt] |= 1 << bit
}

func (this *MyHashSet) Remove(key int) {
	byt := key / 8
	bit := key % 8
	this.set[byt] &= ^(1 << bit)
}

func (this *MyHashSet) Contains(key int) bool {
	byt := key / 8
	bit := key % 8
	v := this.set[byt]
	return v&(1<<bit) != 0
}
