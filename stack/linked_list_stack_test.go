package stack

import "testing"

func TestLinkedListStack(t *testing.T) {
	const size = 100000
	stack, err := NewLinkedListStack[int]()

	if err != nil {
		t.Error("creating new linked list stack", err)
	}

	for i := 0; i < size; i++ {
		if err := stack.Push(i); err != nil {
			t.Error("pushing value to stack", err, "at", i)
		}
	}

	if stack.isEmpty() {
		t.Error("expected stack not to be empty")
	}

	if stack.Peek() != size-1 {
		t.Error("expected head value to be", size-1, "but got", stack.Peek())
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
