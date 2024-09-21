package queue

type ArrayQueue[T any] struct {
	elements []T
	size     int
	tail     int
	head     int
}

func (queue *ArrayQueue[T]) Enqueue(value T) {
	if !queue.isFull() {
		queue.elements = append(queue.elements, value)
		queue.tail++
	}
}

func (queue *ArrayQueue[T]) Dequeue() (T, bool) {
	if !queue.IsEmpty() {
		value := queue.elements[queue.head]
		queue.head++
		return value, true
	}

	var value T
	return value, false
}

func (queue *ArrayQueue[T]) Peek() (T, bool) {
	if !queue.IsEmpty() {
		return queue.elements[queue.head], true
	}

	var value T
	return value, false
}

func (queue *ArrayQueue[T]) IsEmpty() bool {
	return queue.tail == queue.head
}

func (queue *ArrayQueue[T]) Size() int {
	return queue.size
}

func (queue *ArrayQueue[T]) isFull() bool {
	return queue.size == queue.tail
}

func NewArrayQueue[T any](size int) Queue[T] {
	return &ArrayQueue[T]{elements: make([]T, 0, size), size: size}
}
