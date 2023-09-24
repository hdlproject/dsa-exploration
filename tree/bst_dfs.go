package tree

import "fmt"

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

	if curr.Left != nil {
		res := b.search(curr.Left, data)
		if res != nil {
			return res
		}
	}

	if curr.Right != nil {
		res := b.search(curr.Right, data)
		if res != nil {
			return res
		}
	}

	return nil
}

func (b *BSTDFS) Traverse() {
	b.traverse(b.root)
	fmt.Println()
}

func (b *BSTDFS) traverse(curr *BST) {
	fmt.Printf("%d ", curr.data)

	b.curr = curr

	if curr.Left != nil {
		b.traverse(curr.Left)
	}

	if curr.Right != nil {
		b.traverse(curr.Right)
	}
}
