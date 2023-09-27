package tree

import (
	"fmt"
	"math"
)

type BalancedBST struct {
	data          int
	parent        *BalancedBST
	left          *BalancedBST
	right         *BalancedBST
	balanceFactor int
}

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
}

func (b *BalancedBST) IsBalanced() bool {
	var height int
	return b.isBalanced(b, &height)
}

func (b *BalancedBST) isBalanced(curr *BalancedBST, height *int) bool {
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

func (b *BalancedBST) GetImbalancedNode() *BalancedBST {
	var height int
	return b.getImbalancedNode(b, &height)
}

func (b *BalancedBST) getImbalancedNode(curr *BalancedBST, height *int) *BalancedBST {
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

func (b *BalancedBST) RotateLeft() {
	if b.parent != nil {
		if b.parent.right == b {
			b.parent.right = b.right
		} else if b.parent.left == b {
			b.parent.left = b.right
		}

		b.right.parent = b.parent
	}

	b.parent = b.right

	if b.parent.left != nil {
		b.parent.left.parent = b
	}
	b.right = b.parent.left

	b.parent.left = b
}

func (b *BalancedBST) RotateRight() {
	if b.parent != nil {
		if b.parent.left == b {
			b.parent.left = b.left
		} else if b.parent.right == b {
			b.parent.right = b.left
		}

		b.left.parent = b.parent
	}

	b.parent = b.left

	if b.parent.right != nil {
		b.parent.right.parent = b
	}
	b.left = b.parent.right

	b.parent.right = b
}
