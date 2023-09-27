package tree

import (
	"fmt"
	"strings"
)

type BSTDFS struct {
	root BSTTraversal
	curr BSTTraversal
}

func CreateBSTDFS(bst BSTTraversal) *BSTDFS {
	return &BSTDFS{
		root: bst,
		curr: bst,
	}
}

func (b *BSTDFS) Search(data int) BSTTraversal {
	return b.search(b.root, data)
}

func (b *BSTDFS) search(curr BSTTraversal, data int) BSTTraversal {
	b.curr = curr

	if curr.GetData() == data {
		return curr
	}

	if curr.GetLeft() != nil {
		res := b.search(curr.GetLeft(), data)
		if res != nil {
			return res
		}
	}

	if curr.GetRight() != nil {
		res := b.search(curr.GetRight(), data)
		if res != nil {
			return res
		}
	}

	return nil
}

func (b *BSTDFS) Traverse() string {
	var result string
	b.traverse(b.root, &result)

	return strings.Trim(result, " ")
}

func (b *BSTDFS) traverse(curr BSTTraversal, result *string) {
	res := fmt.Sprintf("%s %d", *result, curr.GetData())

	*result = res

	b.curr = curr

	if curr.GetLeft() != nil {
		b.traverse(curr.GetLeft(), result)
	}

	if curr.GetRight() != nil {
		b.traverse(curr.GetRight(), result)
	}
}
