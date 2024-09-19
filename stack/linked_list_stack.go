package stack

import (
	"log"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value, next: nil}
}

type linkedListStack[T any] struct {
	head *Node[T]
}

func (stack *linkedListStack[T]) Push(value T) error {
	newNode := NewNode(value)

	if stack.head == nil {
		stack.head = newNode
	} else {
		newNode.next = stack.head
		stack.head = newNode
	}

	return nil
}

func (stack *linkedListStack[T]) Pop() T {
	if stack.isEmpty() {
		log.Fatal("stack is empty")
	}

	poppedValue := stack.head.value
	stack.head = stack.head.next
	return poppedValue
}

func (stack *linkedListStack[T]) Peek() T {
	return stack.head.value
}

func (stack *linkedListStack[T]) isEmpty() bool {
	return stack.head == nil
}

func NewLinkedListStack[T any]() (Stack[T], error) {
	return &linkedListStack[T]{head: nil}, nil
}
