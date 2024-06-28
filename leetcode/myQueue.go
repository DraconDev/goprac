package leetcode

type MyQueue struct {
	stack   []int
	reverse []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)

}

func (this *MyQueue) Pop() int {
	if len(this.reverse) == 0 {
		this.transfer()
	}

	if len(this.reverse) == 0 {
		return 0
	}
	result := this.reverse[len(this.reverse)-1]
	this.reverse = this.reverse[:len(this.reverse)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.reverse) == 0 {
		this.transfer()
	}

	if len(this.reverse) == 0 {
		return 0
	}

	return this.reverse[len(this.reverse)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack) == 0 && len(this.reverse) == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyQueue) transfer() {
	for len(this.stack) > 0 {
		this.reverse = append(this.reverse, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
}
