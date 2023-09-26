package tree

import (
	"testing"
)

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
func TestBalancedBST_GetImbalancedNode(t *testing.T) {
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

	imbalancedNode := balancedBST.GetImbalancedNode()
	if imbalancedNode == nil {
		t.Fatalf("expect not nil imbalanced_node")
	}

	parentNode.Insert(CreateBST(7))
	parentNode.Insert(CreateBST(9))
	parentNode.Insert(CreateBST(6))

	isBalanced = balancedBST.IsBalanced()
	if !isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", true, isBalanced)
	}

	imbalancedNode = balancedBST.GetImbalancedNode()
	if imbalancedNode != nil {
		t.Fatalf("expect nil imbalanced_node")
	}
}

/*
	 imbalanced tree
				5
		  3			  8
	 	2	4
	  1
*/
/*
	 after rotating
				3
		  2			  5
	 	1			4	8
*/
func TestBalancedBST_RotateRight(t *testing.T) {
	parentNode := CreateBST(5)
	parentNodeAfterRotation := CreateBST(3)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBST(8))
	parentNode.Insert(CreateBST(2))
	parentNode.Insert(CreateBST(4))
	parentNode.Insert(CreateBST(1))

	balancedBST := CreateBalancedBST(parentNode)
	balancedBST.RotateRight(parentNode)

	bstDFS := CreateBSTDFS(parentNodeAfterRotation)
	res := bstDFS.Traverse()

	expectedRes := "3 2 1 5 4 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}

/*
	 imbalanced tree
				2
		  1			  4
					3	5
						  8
*/
/*
	 after rotating
				4
		  2			  5
	 	1	3			8
*/
func TestBalancedBST_RotateLeft(t *testing.T) {
	parentNode := CreateBST(2)
	parentNodeAfterRotation := CreateBST(4)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBST(1))
	parentNode.Insert(CreateBST(5))
	parentNode.Insert(CreateBST(3))
	parentNode.Insert(CreateBST(8))

	balancedBST := CreateBalancedBST(parentNode)
	balancedBST.RotateLeft(parentNode)

	bstDFS := CreateBSTDFS(parentNodeAfterRotation)
	res := bstDFS.Traverse()

	expectedRes := "4 2 1 3 5 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}
