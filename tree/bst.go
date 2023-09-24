package tree

import "fmt"

type BST struct {
	data  int
	Left  *BST
	Right *BST
}

type nodeType int

const (
	left nodeType = iota
	right
)

func CreateBST(data int) *BST {
	return &BST{
		data: data,
	}
}

func (b *BST) Insert(nodeType nodeType, node *BST) {
	switch nodeType {
	case left:
		b.Left = node
	case right:
		b.Right = node
	}
}

func (b *BST) String() string {
	return fmt.Sprintf("%d", b.data)
}
