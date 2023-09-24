package tree

import (
	"fmt"

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

		if b.curr.Left != nil {
			b.queue.Enqueue(b.curr.Left)
		}

		if b.curr.Right != nil {
			b.queue.Enqueue(b.curr.Right)
		}
	}

	return nil
}

func (b *BSTBFS) Traverse() {
	b.queue.Enqueue(b.root)

	for !b.queue.IsEmpty() {
		b.currQ = b.queue.Dequeue()

		b.curr = b.currQ.(*BST)
		fmt.Printf("%s ", b.curr)

		if b.curr.Left != nil {
			b.queue.Enqueue(b.curr.Left)
		}

		if b.curr.Right != nil {
			b.queue.Enqueue(b.curr.Right)
		}
	}

	fmt.Println()
}
