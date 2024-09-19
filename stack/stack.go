package stack

type Stack[T any] interface {
	Push(T) error
	Pop() T
	Peek() T
	isEmpty() bool
}
