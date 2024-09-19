package stack

import "testing"

func TestArrayStack(t *testing.T) {
	const size = 100000
	stack, err := NewArrayStack[int](size)

	if err != nil {
		t.Error("creating new array stack", err)
	}

	for i := 0; i < size; i++ {
		if err := stack.Push(i); err != nil {
			t.Error("pushing value to stack", err, "at", i)
		}
	}

	if stack.isEmpty() {
		t.Error("expected stack not to be empty")
	}

	for i := size - 1; i >= 0; i-- {
		if popped := stack.Pop(); popped != i {
			t.Error("expected popped value to match", i, "but got", popped)
		}
	}

	if !stack.isEmpty() {
		t.Error("expected stack to be empty")
	}
}
