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

	imbalancedNode, rt := b.updateBalanceFactor(node)

	switch rt {
	case leftRotationType:
		b.rotateLeft(imbalancedNode)
	case rightRotationType:
		b.rotateRight(imbalancedNode)
	case leftRightRotationType:
	case rightLeftRotationType:
	}

}

func (b *BalancedBST) updateBalanceFactor(node *BalancedBST) (*BalancedBST, rotationType) {
	var backTrackingRecord string

	curr := node
	rt := noRotationType
	for curr.parent != nil {
		if curr.parent.left == curr {
			curr.parent.balanceFactor++
			backTrackingRecord = fmt.Sprintf("%s %d", backTrackingRecord, 1)
		} else if curr.parent.right == curr {
			curr.parent.balanceFactor--
			backTrackingRecord = fmt.Sprintf("%s %d", backTrackingRecord, -1)
		}

		curr = curr.parent

		if math.Abs(float64(curr.balanceFactor)) > 1 {
			if node.data < curr.left.data {
				rt = rightRotationType
			} else if node.data > curr.right.data {
				rt = leftRotationType
			} else if node.data > curr.left.data {
				rt = leftRightRotationType
			} else if node.data < curr.right.data {
				rt = rightLeftRotationType
			}

			return curr, rt
		}
	}

	return nil, rt
}

func (b *BalancedBST) RotateLeft() {
	b.rotateLeft(b)
}

func (b *BalancedBST) rotateLeft(curr *BalancedBST) {
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

	curr.updateBalanceFactorAfterRotation()
}

func (b *BalancedBST) RotateRight() {
	b.rotateRight(b)
}

func (b *BalancedBST) rotateRight(curr *BalancedBST) {
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

	curr.updateBalanceFactorAfterRotation()
}

func (b *BalancedBST) rotateLeftRight(curr *BalancedBST) {

}

func (b *BalancedBST) rotateRightLeft(curr *BalancedBST) {

}

func (b *BalancedBST) updateBalanceFactorAfterRotation() {
	b.parent.balanceFactor = 0

	b.balanceFactor = 0
	if b.left != nil {
		b.balanceFactor += 1
	}
	if b.right != nil {
		b.balanceFactor -= 1
	}
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
