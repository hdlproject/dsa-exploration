package tree

import "testing"

/*
				1
		  2			  3
	 	4	5		6	7
			  8	  9
*/
func TestBSTDFS(t *testing.T) {
	parentNode := CreateBST(1)
	leftNode := CreateBST(2)
	rightNode := CreateBST(3)
	parentNode.Insert(left, leftNode)
	parentNode.Insert(right, rightNode)

	leftLeftNode := CreateBST(4)
	leftRightNode := CreateBST(5)
	leftNode.Insert(left, leftLeftNode)
	leftNode.Insert(right, leftRightNode)

	leftRightRightNode := CreateBST(8)
	leftRightNode.Insert(right, leftRightRightNode)

	rightLeftNode := CreateBST(6)
	rightRightNode := CreateBST(7)
	rightNode.Insert(left, rightLeftNode)
	rightNode.Insert(right, rightRightNode)

	rightLeftLeftNode := CreateBST(9)
	rightLeftNode.Insert(left, rightLeftLeftNode)

	// expected --->> 1 2 4 5 8 3 6 9 7
	bstDFS := CreateBSTDFS(parentNode)
	bstDFS.Traverse()

	expectedData := 5
	res := bstDFS.Search(expectedData)
	if res == nil {
		t.Fatalf("expect not nil result")
	}
	if res.data != expectedData {
		t.Fatalf("expect data %d but got %d", expectedData, res.data)
	}

	//// expected --->> 1 2 3 4 5 6 7 8 9
	bstBFS := CreateBSTBFS(parentNode)
	bstBFS.Traverse()

	res = bstBFS.Search(expectedData)
	if res == nil {
		t.Fatalf("expect not nil result")
	}
	if res.data != expectedData {
		t.Fatalf("expect data %d but got %d", expectedData, res.data)
	}
}
