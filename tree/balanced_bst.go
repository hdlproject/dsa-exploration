package tree

import (
	"math"
)

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
	if curr == nil {
		return true
	}

	var leftHeight, rightHeight int
	isLeftBalanced := b.isBalanced(curr.left, &leftHeight)
	isRightBalanced := b.isBalanced(curr.right, &rightHeight)

	*height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1

	if math.Abs(float64(leftHeight-rightHeight)) <= 1 {
		return isLeftBalanced && isRightBalanced
	}

	return false
}

func (b *BalancedBST) GetImbalancedNode() *BST {
	var height int
	return b.getImbalancedNode(b.root, &height)
}

func (b *BalancedBST) getImbalancedNode(curr *BST, height *int) *BST {
	if curr == nil {
		return nil
	}

	var leftHeight, rightHeight int
	imbalancedLeftNode := b.getImbalancedNode(curr.left, &leftHeight)
	if imbalancedLeftNode != nil {
		return curr.left
	}
	imbalancedRightNode := b.getImbalancedNode(curr.right, &rightHeight)
	if imbalancedRightNode != nil {
		return curr.right
	}

	*height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1

	if math.Abs(float64(leftHeight-rightHeight)) <= 1 {
		return nil
	}

	return curr
}

func (b *BalancedBST) RotateLeft(curr *BST) {
	if curr.parent != nil {
		if curr.parent.right == curr {
			curr.parent.right = curr.right
		} else if curr.parent.left == curr {
			curr.parent.left = curr.right
		}

		curr.right.parent = curr.parent
	}

	curr.parent = curr.right

	if curr.parent.left != nil {
		curr.parent.left.parent = curr
	}
	curr.right = curr.parent.left

	curr.parent.left = curr
}

func (b *BalancedBST) RotateRight(curr *BST) {
	if curr.parent != nil {
		if curr.parent.left == curr {
			curr.parent.left = curr.left
		} else if curr.parent.right == curr {
			curr.parent.right = curr.left
		}

		curr.left.parent = curr.parent
	}

	curr.parent = curr.left

	if curr.parent.right != nil {
		curr.parent.right.parent = curr
	}
	curr.left = curr.parent.right

	curr.parent.right = curr
}
