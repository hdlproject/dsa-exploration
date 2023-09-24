package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := CreateQueue()
	for i := 0; i < 3; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < 3; i++ {
		node := queue.Dequeue()
		if node.(int) != i {
			t.Fatalf("expect node %d but got %d", i, node)
		}
	}
}
