package node

import "fmt"

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func CreateNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

func (n *Node) SetPrev(node *Node) {
	n.prev = node
}

func (n *Node) GetData() interface{} {
	return n.data
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}
