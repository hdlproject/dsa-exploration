package tree

import "testing"

/* imbalanced tree
			5
	  2			  8
 	1	3
		  4
*/
/* balanced tree
			5
	  2			  8
 	1	3		7	9
		  4	  6
*/
func TestBalancedBST(t *testing.T) {
	parentNode := CreateBST(5)
	parentNode.Insert(CreateBST(2))
	parentNode.Insert(CreateBST(8))
	parentNode.Insert(CreateBST(1))
	parentNode.Insert(CreateBST(3))
	parentNode.Insert(CreateBST(4))

	balancedBST := CreateBalancedBST(parentNode)
	isBalanced := balancedBST.IsBalanced()
	if isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", false, isBalanced)
	}

	parentNode.Insert(CreateBST(7))
	parentNode.Insert(CreateBST(9))
	parentNode.Insert(CreateBST(6))

	isBalanced = balancedBST.IsBalanced()
	if !isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", true, isBalanced)
	}
}
