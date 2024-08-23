package queue

import "dsa-exploration/node"

type Queue struct {
	firstNode *node.Node
	lastNode  *node.Node
}

func CreateQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data interface{}) {
	n := node.CreateNode(data)

	if q.lastNode == nil {
		q.lastNode = n
		q.firstNode = n
	} else {
		q.lastNode.SetNext(n)
		q.lastNode = n
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.firstNode == nil {
		return nil
	}

	currNode := q.firstNode
	q.firstNode = q.firstNode.GetNext()

	if q.firstNode == nil {
		q.lastNode = nil
	}

	return currNode.GetData()
}

func (q *Queue) IsEmpty() bool {
	return q.firstNode == nil
}
