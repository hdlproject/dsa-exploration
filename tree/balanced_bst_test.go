package tree

import "testing"

/* imbalanced tree
			1
	  2			  3
 	4	5
		  6
*/
/* balanced tree
			1
	  2			  3
 	4	5			7
		  6
*/
func TestBalancedBST(t *testing.T) {
	parentNode := CreateBST(1)
	leftNode := CreateBST(2)
	rightNode := CreateBST(3)
	parentNode.Insert(left, leftNode)
	parentNode.Insert(right, rightNode)

	leftLeftNode := CreateBST(4)
	leftRightNode := CreateBST(5)
	leftNode.Insert(left, leftLeftNode)
	leftNode.Insert(right, leftRightNode)

	leftRightRightNode := CreateBST(6)
	leftRightNode.Insert(right, leftRightRightNode)

	balancedBST := CreateBalancedBST(parentNode)
	isBalanced := balancedBST.IsBalanced()
	if isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", false, isBalanced)
	}

	rightRightNode := CreateBST(7)
	rightNode.Insert(right, rightRightNode)

	isBalanced = balancedBST.IsBalanced()
	if !isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", true, isBalanced)
	}
}
