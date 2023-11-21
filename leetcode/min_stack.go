package leetcode

func TestMinStackConstructor() {
	s := MinStack{}
	s.GetMin()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.GetMin()
	s.Pop()
	s.Top()
	s.GetMin()

}

type MinStack struct {
	stack    []int
	minStack []int
}

func MinStackConstructor() MinStack {
	s := MinStack{}

	return s
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minStack) == 0 || val <= this.GetMin() {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {

	if this.Top() == this.GetMin() {
		this.minStack = this.minStack[:len(this.stack)-1]
		// this.minStack = append(this.minStack[:len(this.minStack)-1])
	}
	this.stack = append(this.stack[0 : len(this.stack)-1])

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}
