package queue

import "fmt"

//type NodeDataInf interface {
//	String() string
//}
//
//type SimpleNode struct {
//	data int
//}
//
//func (s *SimpleNode) String() string {
//	return fmt.Sprintf("%d", s.data)
//}
//
//func CreateSimpleNode(data int) *SimpleNode {
//	return &SimpleNode{
//		data: data,
//	}
//}

type Node struct {
	data interface{}
	next *Node
}

func CreateNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n)
}

type Queue struct {
	firstNode *Node
	lastNode  *Node
}

func CreateQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data interface{}) {
	node := CreateNode(data)

	if q.lastNode == nil {
		q.lastNode = node
		q.firstNode = node
	} else {
		q.lastNode.next = node
		q.lastNode = node
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.firstNode == nil {
		return nil
	}

	currNode := q.firstNode
	q.firstNode = q.firstNode.next

	if q.firstNode == nil {
		q.lastNode = nil
	}

	return currNode.data
}

func (q *Queue) IsEmpty() bool {
	return q.firstNode == nil
}
