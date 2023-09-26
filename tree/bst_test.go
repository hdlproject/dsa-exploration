package tree

import (
	"testing"
)

/*
				5
		  2			  8
	 	1	3		7	9
			  4	  6
*/
func TestBST(t *testing.T) {
	parentNode := CreateBST(5)
	parentNode.Insert(CreateBST(2))
	parentNode.Insert(CreateBST(8))
	parentNode.Insert(CreateBST(1))
	parentNode.Insert(CreateBST(3))
	parentNode.Insert(CreateBST(7))
	parentNode.Insert(CreateBST(9))
	parentNode.Insert(CreateBST(4))
	parentNode.Insert(CreateBST(6))

	expectedRes := "5 2 1 3 4 8 7 6 9"
	bstDFS := CreateBSTDFS(parentNode)
	res := bstDFS.Traverse()
	if expectedRes != res {
		t.Fatalf("expect result %s but got %s", expectedRes, res)
	}

	expectedData := 5
	searchRes := bstDFS.Search(expectedData)
	if searchRes == nil {
		t.Fatalf("expect not nil result")
	}
	if searchRes.data != expectedData {
		t.Fatalf("expect data %d but got %d", expectedData, searchRes.data)
	}

	expectedRes = "5 2 8 1 3 7 9 4 6"
	bstBFS := CreateBSTBFS(parentNode)
	res = bstBFS.Traverse()
	if expectedRes != res {
		t.Fatalf("expect result %s but got %s", expectedRes, res)
	}

	searchRes = bstBFS.Search(expectedData)
	if searchRes == nil {
		t.Fatalf("expect not nil result")
	}
	if searchRes.data != expectedData {
		t.Fatalf("expect data %d but got %d", expectedData, searchRes.data)
	}
}
