package stack

import (
	"errors"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEpmpty() bool {
	if len(stack) > 0 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Can't Top() an empty stack!")
	}
	return stack[len(stack)-1], nil
}