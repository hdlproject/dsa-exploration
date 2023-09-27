package tree

type BSTTraversal interface {
	GetLeft() BSTTraversal
	GetRight() BSTTraversal
	GetData() int
}
