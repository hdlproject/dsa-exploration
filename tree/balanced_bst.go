package tree

import "math"

type BalancedBST struct {
	root *BST
}

func CreateBalancedBST(bst *BST) *BalancedBST {
	return &BalancedBST{
		root: bst,
	}
}

func (b *BalancedBST) IsBalanced() bool {
	var height int
	return b.isBalanced(b.root, &height)
}

func (b *BalancedBST) isBalanced(curr *BST, height *int) bool {
	var leftHeight, rightHeight int

	if curr == nil {
		return true
	}

	isLeftBalanced := b.isBalanced(curr.Left, &leftHeight)
	isRightBalanced := b.isBalanced(curr.Right, &rightHeight)

	*height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1

	if math.Abs(float64(leftHeight-rightHeight)) <= 1 {
		return isLeftBalanced && isRightBalanced
	}

	return false
}
