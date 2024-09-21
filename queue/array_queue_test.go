package queue

import "testing"

func TestNewArrayQueue(t *testing.T) {
	const size = 100000
	queue := NewArrayQueue[int](size)

	for i := 0; i < size; i++ {
		queue.Enqueue(i)
	}

	if queue.IsEmpty() {
		t.Errorf("expected queue not to be empty")
	}

	for i := 0; i < size; i++ {
		peek, ok := queue.Peek()
		if !ok {
			t.Fatal("expected peeked value")
		}

		value, ok := queue.Dequeue()
		if !ok {
			t.Fatal("expected dequeued value")
		}

		if value != peek {
			t.Error("expected", peek, "got", value)
		}

		if value != i {
			t.Error("expected", i, "got", value)
		}

	}

	if !queue.IsEmpty() {
		t.Errorf("expected queue to be empty")
	}
}
