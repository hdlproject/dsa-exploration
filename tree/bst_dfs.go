package tree

import (
	"fmt"
	"strings"
)

type BSTDFS struct {
	root *BST
	curr *BST
}

func CreateBSTDFS(bst *BST) *BSTDFS {
	return &BSTDFS{
		root: bst,
		curr: bst,
	}
}

func (b *BSTDFS) Search(data int) *BST {
	return b.search(b.root, data)
}

func (b *BSTDFS) search(curr *BST, data int) *BST {
	b.curr = curr

	if curr.data == data {
		return curr
	}

	if curr.left != nil {
		res := b.search(curr.left, data)
		if res != nil {
			return res
		}
	}

	if curr.right != nil {
		res := b.search(curr.right, data)
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

func (b *BSTDFS) traverse(curr *BST, result *string) {
	res := fmt.Sprintf("%s %d", *result, curr.data)

	*result = res

	b.curr = curr

	if curr.left != nil {
		b.traverse(curr.left, result)
	}

	if curr.right != nil {
		b.traverse(curr.right, result)
	}
}
