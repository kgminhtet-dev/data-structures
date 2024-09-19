package stack

import "fmt"

type arrayStack[T any] struct {
	values          []T
	length          int
	currentPosition int
}

func (stack *arrayStack[T]) isFull() bool {
	return stack.currentPosition == stack.length-1
}

func (stack *arrayStack[T]) Push(value T) error {
	if stack.isFull() {
		return fmt.Errorf("stack is full")
	}
	stack.values = append(stack.values, value)
	stack.currentPosition++
	return nil
}

func (stack *arrayStack[T]) Pop() T {
	popValue := stack.values[stack.currentPosition]
	stack.values = stack.values[:stack.currentPosition]
	stack.currentPosition--
	return popValue
}

func (stack *arrayStack[T]) Peek() T {
	return stack.values[stack.currentPosition]
}

func (stack *arrayStack[T]) isEmpty() bool {
	return stack.currentPosition == -1
}

func NewArrayStack[T any](size int) (Stack[T], error) {
	if size <= 0 {
		return nil, fmt.Errorf("stack size must be greater than zero")
	}
	return &arrayStack[T]{
		values:          make([]T, 0, size),
		length:          size,
		currentPosition: -1,
	}, nil
}
