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
	parentNode := CreateBalancedBST(5)
	parentNode.Insert(CreateBalancedBST(2))
	parentNode.Insert(CreateBalancedBST(8))
	parentNode.Insert(CreateBalancedBST(1))
	parentNode.Insert(CreateBalancedBST(3))
	parentNode.Insert(CreateBalancedBST(4))

	isBalanced := parentNode.IsBalanced()
	if isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", false, isBalanced)
	}

	imbalancedNode := parentNode.GetImbalancedNode()
	if imbalancedNode == nil {
		t.Fatalf("expect not nil imbalanced_node")
	}

	parentNode.Insert(CreateBalancedBST(7))
	parentNode.Insert(CreateBalancedBST(9))
	parentNode.Insert(CreateBalancedBST(6))

	isBalanced = parentNode.IsBalanced()
	if !isBalanced {
		t.Fatalf("expect is_balanced %v but got %v", true, isBalanced)
	}

	imbalancedNode = parentNode.GetImbalancedNode()
	if imbalancedNode != nil {
		t.Fatalf("expect nil imbalanced_node")
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
	parentNode := CreateBalancedBST(2)
	parentNodeAfterRotation := CreateBalancedBST(4)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBalancedBST(1))
	parentNode.Insert(CreateBalancedBST(5))
	parentNode.Insert(CreateBalancedBST(3))
	parentNode.Insert(CreateBalancedBST(8))

	parentNode.RotateLeft()

	bstDFS := CreateBSTDFS(parentNodeAfterRotation)
	res := bstDFS.Traverse()

	expectedRes := "4 2 1 3 5 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
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
	parentNode := CreateBalancedBST(5)
	parentNodeAfterRotation := CreateBalancedBST(3)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBalancedBST(8))
	parentNode.Insert(CreateBalancedBST(2))
	parentNode.Insert(CreateBalancedBST(4))
	parentNode.Insert(CreateBalancedBST(1))

	parentNode.RotateRight()

	bstDFS := CreateBSTDFS(parentNodeAfterRotation)
	res := bstDFS.Traverse()

	expectedRes := "3 2 1 5 4 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}

/*
	 imbalanced tree
				5
		  2			  8
	 	1	3
			  4
*/
/*
	 after first rotating
				5
		  3			  8
	 	2	4
	  1
*/
/*
	 after second rotating
				3
		  2			  5
	 	1			4	8
*/
func TestBalancedBST_RotateLeftRight(t *testing.T) {
	parentNode := CreateBalancedBST(5)
	firstRotationNode := CreateBalancedBST(2)
	parentNode.Insert(firstRotationNode)
	parentNodeAfterRotation := CreateBalancedBST(3)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBalancedBST(4))
	parentNode.Insert(CreateBalancedBST(1))
	parentNode.Insert(CreateBalancedBST(8))

	firstRotationNode.RotateLeft()

	bstDFS := CreateBSTDFS(parentNode)
	res := bstDFS.Traverse()

	expectedRes := "5 3 2 1 4 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}

	parentNode.RotateRight()

	bstDFS = CreateBSTDFS(parentNodeAfterRotation)
	res = bstDFS.Traverse()

	expectedRes = "3 2 1 5 4 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}

/*
	 imbalanced tree
				2
		  1			  5
					4	6
				  3
*/
/*
	 after first rotating
				2
		  1			  4
					3	5
						  6
*/
/*
	 after second rotating
				3
		  2			  5
	 	1			4	8
*/
func TestBalancedBST_RotateRightLeft(t *testing.T) {
	parentNode := CreateBalancedBST(5)
	firstRotationNode := CreateBalancedBST(2)
	parentNode.Insert(firstRotationNode)
	parentNodeAfterRotation := CreateBalancedBST(3)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBalancedBST(4))
	parentNode.Insert(CreateBalancedBST(1))
	parentNode.Insert(CreateBalancedBST(8))

	firstRotationNode.RotateLeft()

	bstDFS := CreateBSTDFS(parentNode)
	res := bstDFS.Traverse()

	expectedRes := "5 3 2 1 4 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}

	parentNode.RotateRight()

	bstDFS = CreateBSTDFS(parentNodeAfterRotation)
	res = bstDFS.Traverse()

	expectedRes = "3 2 1 5 4 8"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}
