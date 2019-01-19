// stack.go
// Author: hyan23
// Date: 2018.11.01

package main

type Stack struct {
	sta []interface{}
	capacity uint
	top uint
}

func (stack *Stack) Init (capacity uint) {
	stack.sta = make([]interface{}, capacity)
	stack.capacity = uint (len(stack.sta))
	stack.top = stack.capacity
}

func (stack *Stack) Capacity() uint {
	return stack.capacity
}

func (stack *Stack) Empty() bool {
	return stack.top == stack.capacity
}

func (stack *Stack) Full() bool {
	return stack.top == 0
}

func (stack *Stack) Size() uint {
	return stack.capacity - stack.top
}

func (stack *Stack) Push (e interface{}) bool {
	if stack.top > 0 {
		stack.top --
		stack.sta[stack.top] = e
		return true
	} else {
		return false
	}
}

func (stack *Stack) Top() (interface{}, bool) {
	if stack.top == stack.capacity {
		return nil, false
	} else {
		return stack.sta[stack.top], true
	}
}

func (stack *Stack) Pop() (interface{}, bool) {
	defer func () {
		if stack.top < stack.capacity {
			stack.top ++
		}
	}()
	return stack.Top()
}