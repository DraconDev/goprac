package leetcode

type MinStack struct {
	stack []int
}

func MinStackConstructor() MinStack {
	s := MinStack{}
	// s.Push(3)
	// s.Push(4)
	// s.Push(2)
	// s.Push(8)
	// s.GetMin()
	// s.Pop()
	// s.GetMin()
	// s.Pop()
	// s.GetMin()
	// s.Pop()
	// s.GetMin()
	// s.Pop()
	// s.GetMin()

	// s.Top()
	return s
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = append(this.stack[0 : len(this.stack)-1])

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	min := this.stack[0]
	for _, v := range this.stack {
		if v < min {
			min = v
		}
	}
	return min
}
