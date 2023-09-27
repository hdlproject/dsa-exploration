package tree

import "fmt"

type BST struct {
	data   int
	parent *BST
	left   *BST
	right  *BST
}

func CreateBST(data int) *BST {
	return &BST{
		data: data,
	}
}

func (b *BST) GetLeft() BSTTraversal {
	if b.left == nil {
		return nil
	}
	return b.left
}

func (b *BST) GetRight() BSTTraversal {
	if b.right == nil {
		return nil
	}
	return b.right
}

func (b *BST) GetData() int {
	return b.data
}

func (b *BST) String() string {
	return fmt.Sprintf("%d", b.data)
}

func (b *BST) Insert(node *BST) {
	curr := b

	for {
		var next *BST
		if node.data <= curr.data {
			next = curr.left
		} else {
			next = curr.right
		}

		if next == nil {
			break
		}

		curr = next
	}

	if node.data <= curr.data {
		curr.left = node
	} else {
		curr.right = node
	}
	node.parent = curr
}
