package tree

import (
	"fmt"
	"math"
)

type BalancedBST struct {
	data   int
	parent *BalancedBST
	left   *BalancedBST
	right  *BalancedBST
	height int
}

type rotationType int

const (
	noRotationType rotationType = iota
	leftRotationType
	rightRotationType
	leftRightRotationType
	rightLeftRotationType
)

func CreateBalancedBST(data int) *BalancedBST {
	return &BalancedBST{
		data: data,
	}
}

func (b *BalancedBST) GetLeft() BSTTraversal {
	if b.left == nil {
		return nil
	}
	return b.left
}

func (b *BalancedBST) GetRight() BSTTraversal {
	if b.right == nil {
		return nil
	}
	return b.right
}

func (b *BalancedBST) GetData() int {
	return b.data
}

func (b *BalancedBST) String() string {
	return fmt.Sprintf("%d", b.data)
}

func (b *BalancedBST) Insert(node *BalancedBST) {
	curr := b

	for {
		var next *BalancedBST
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

	for {
		balanceFactor := curr.getBalanceFactor()
		if balanceFactor > 1 {
			if node.data < curr.left.data {
				curr = b.rotateRight(curr)
			} else if node.data > curr.left.data {
				b.rotateLeft(curr.left)
				curr.left.height = int(math.Max(float64(curr.left.left.getHeight()), float64(curr.left.right.getHeight()))) + 1

				curr = b.rotateRight(curr)
			}
		} else if balanceFactor < -1 {
			if node.data > curr.right.data {
				curr = b.rotateLeft(curr)
			} else if node.data < curr.right.data {
				b.rotateRight(curr.right)
				curr.right.height = int(math.Max(float64(curr.right.left.getHeight()), float64(curr.right.right.getHeight()))) + 1

				b.rotateLeft(curr)
			}
		}

		curr.height = int(math.Max(float64(curr.left.getHeight()), float64(curr.right.getHeight()))) + 1

		if curr.parent == nil {
			break
		}

		curr = curr.parent
	}
}

func (b *BalancedBST) getBalanceFactor() int {
	return b.left.getHeight() - b.right.getHeight()
}

func (b *BalancedBST) getHeight() int {
	if b == nil {
		return 0
	}

	return b.height
}

func (b *BalancedBST) RotateLeft() {
	b.rotateLeft(b)
}

func (b *BalancedBST) rotateLeft(curr *BalancedBST) (newStartingNode *BalancedBST) {
	if curr.parent != nil {
		if curr.data > curr.parent.data {
			curr.parent.right = curr.right
		} else if curr.data <= curr.parent.data {
			curr.parent.left = curr.right
		}
	}
	curr.right.parent = curr.parent

	curr.parent = curr.right

	if curr.parent.left != nil {
		curr.parent.left.parent = curr
	}
	curr.right = curr.parent.left

	curr.parent.left = curr

	curr.height = int(math.Max(float64(curr.left.getHeight()), float64(curr.right.getHeight()))) + 1

	return curr.parent
}

func (b *BalancedBST) RotateRight() {
	b.rotateRight(b)
}

func (b *BalancedBST) rotateRight(curr *BalancedBST) (newStartingNode *BalancedBST) {
	if curr.parent != nil {
		if curr.parent.left == curr {
			curr.parent.left = curr.left
		} else if curr.parent.right == curr {
			curr.parent.right = curr.left
		}
	}
	curr.left.parent = curr.parent

	curr.parent = curr.left

	if curr.parent.right != nil {
		curr.parent.right.parent = curr
	}
	curr.left = curr.parent.right

	curr.parent.right = curr

	curr.height = int(math.Max(float64(curr.left.getHeight()), float64(curr.right.getHeight()))) + 1

	return curr.parent
}
