package tree

import (
	"fmt"
	"strings"

	"dsa-exploration/queue"
)

type BSTBFS struct {
	root  *BST
	curr  *BST
	currQ interface{}
	queue *queue.Queue
}

func CreateBSTBFS(bst *BST) *BSTBFS {
	return &BSTBFS{
		root:  bst,
		curr:  bst,
		queue: queue.CreateQueue(),
	}
}

func (b *BSTBFS) Search(data int) *BST {
	b.queue.Enqueue(b.root)

	for !b.queue.IsEmpty() {
		b.currQ = b.queue.Dequeue()

		b.curr = b.currQ.(*BST)
		if b.curr.data == data {
			return b.curr
		}

		if b.curr.left != nil {
			b.queue.Enqueue(b.curr.left)
		}

		if b.curr.right != nil {
			b.queue.Enqueue(b.curr.right)
		}
	}

	return nil
}

func (b *BSTBFS) Traverse() string {
	b.queue.Enqueue(b.root)

	var result string
	for !b.queue.IsEmpty() {
		b.currQ = b.queue.Dequeue()

		b.curr = b.currQ.(*BST)
		result = fmt.Sprintf("%s %d", result, b.curr.data)

		if b.curr.left != nil {
			b.queue.Enqueue(b.curr.left)
		}

		if b.curr.right != nil {
			b.queue.Enqueue(b.curr.right)
		}
	}

	return strings.Trim(result, " ")
}
