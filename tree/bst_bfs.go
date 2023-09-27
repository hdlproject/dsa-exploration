package tree

import (
	"fmt"
	"strings"

	"dsa-exploration/queue"
)

type BSTBFS struct {
	root  BSTTraversal
	curr  BSTTraversal
	currQ interface{}
	queue *queue.Queue
}

func CreateBSTBFS(bst BSTTraversal) *BSTBFS {
	return &BSTBFS{
		root:  bst,
		curr:  bst,
		queue: queue.CreateQueue(),
	}
}

func (b *BSTBFS) Search(data int) BSTTraversal {
	b.queue.Enqueue(b.root)

	for !b.queue.IsEmpty() {
		b.currQ = b.queue.Dequeue()

		b.curr = b.currQ.(BSTTraversal)
		if b.curr.GetData() == data {
			return b.curr
		}

		if b.curr.GetLeft() != nil {
			b.queue.Enqueue(b.curr.GetLeft())
		}

		if b.curr.GetRight() != nil {
			b.queue.Enqueue(b.curr.GetRight())
		}
	}

	return nil
}

func (b *BSTBFS) Traverse() string {
	b.queue.Enqueue(b.root)

	var result string
	for !b.queue.IsEmpty() {
		b.currQ = b.queue.Dequeue()

		b.curr = b.currQ.(BSTTraversal)
		result = fmt.Sprintf("%s %d", result, b.curr.GetData())

		if b.curr.GetLeft() != nil {
			b.queue.Enqueue(b.curr.GetLeft())
		}

		if b.curr.GetRight() != nil {
			b.queue.Enqueue(b.curr.GetRight())
		}
	}

	return strings.Trim(result, " ")
}
