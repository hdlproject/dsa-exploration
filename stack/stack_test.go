package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := CreateStack()
	for i := 0; i < 3; i++ {
		stack.Push(i)
	}

	for i := 2; i >= 0; i-- {
		node := stack.Pop()
		if node.(int) != i {
			t.Fatalf("expect node %d but got %d", i, node)
		}
	}
}
