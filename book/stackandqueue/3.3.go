package stackandqueue

type DinnerPlates struct {
	capacity int
	stacks   []Stack
}

func Constructor2(capacity int) DinnerPlates {
	return DinnerPlates{
		capacity: capacity,
		stacks:   []Stack{},
	}
}

func (this *DinnerPlates) Push(val int) {
	lastStack := Stack{}
	if len(this.stacks) > 0 {
		lastStack = this.stacks[len(this.stacks)-1]
	}

	if lastStack.Size() == this.capacity {
		newStack := Stack{}
		newStack.Push(val)
		this.stacks = append(this.stacks, newStack)
	} else {
		lastStack.Push(val)
	}
}

func (this *DinnerPlates) Pop() int {
	lastStack := Stack{}
	if len(this.stacks) > 0 {
		lastStack = this.stacks[len(this.stacks)-1]
	}

	val, _ := lastStack.Pop()
	return val.(int)
}

func (this *DinnerPlates) PopAtStack(index int) int {
	lastStack := Stack{}
	if len(this.stacks) > 0 {
		lastStack = this.stacks[index]
	}

	val, _ := lastStack.Pop()
	return val.(int)
}
