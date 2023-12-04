package stackandqueue

import "fmt"

type MinStack struct {
	elements []int
	min      []int
}

func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		min:      []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)

	if len(this.elements) == 1 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	el := this.elements[len(this.elements)-1]

	if el == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	this.elements = this.elements[:len(this.elements)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func Call3_2() {
	fmt.Println("Validate Solution at https://leetcode.com/problems/min-stack/")
}
