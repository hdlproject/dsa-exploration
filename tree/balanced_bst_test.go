package tree

import (
	"testing"
)

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
	parentNode.Insert(CreateBalancedBST(2))
	parentNode.Insert(CreateBalancedBST(8))
	parentNode.Insert(CreateBalancedBST(1))
	parentNodeAfterRotation := CreateBalancedBST(3)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBalancedBST(4))

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
				4
		  2			  5
	 	1	3			6
*/
func TestBalancedBST_RotateRightLeft(t *testing.T) {
	parentNode := CreateBalancedBST(2)
	parentNode.Insert(CreateBalancedBST(1))
	parentNode.Insert(CreateBalancedBST(5))
	parentNodeAfterRotation := CreateBalancedBST(4)
	parentNode.Insert(parentNodeAfterRotation)
	parentNode.Insert(CreateBalancedBST(6))
	parentNode.Insert(CreateBalancedBST(3))

	bstDFS := CreateBSTDFS(parentNodeAfterRotation)
	res := bstDFS.Traverse()

	expectedRes := "4 2 1 3 5 6"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}

/*
	 initial
				4
		  2			  7
	 	1	3		6	8
				  5		  9
*/
/*
	 after first deletion
				4
		  2			  8
	 	1	3		6	9
				  5
*/
/*
	 after second deletion
				4
		  2			  8
	 	1	3		6	9
				  5
*/
func TestBalancedBST_Delete(t *testing.T) {
	parentNode := CreateBalancedBST(4)
	parentNode.Insert(CreateBalancedBST(2))
	parentNode.Insert(CreateBalancedBST(7))
	parentNode.Insert(CreateBalancedBST(1))
	parentNode.Insert(CreateBalancedBST(3))
	parentNode.Insert(CreateBalancedBST(6))
	parentNode.Insert(CreateBalancedBST(8))
	parentNode.Insert(CreateBalancedBST(5))
	parentNode.Insert(CreateBalancedBST(9))

	parentNode.Delete(CreateBalancedBST(7))

	bstDFS := CreateBSTDFS(parentNode)
	res := bstDFS.Traverse()

	expectedRes := "4 2 1 3 8 6 5 9"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}

	parentNode.Delete(CreateBalancedBST(5))

	res = bstDFS.Traverse()

	expectedRes = "4 2 1 3 8 6 9"
	if res != expectedRes {
		t.Fatalf("expect res %s but got %s", expectedRes, res)
	}
}
