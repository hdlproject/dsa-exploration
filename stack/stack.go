package stack

import "dsa-exploration/node"

type Stack struct {
	firstNode *node.Node
}

func CreateStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data interface{}) {
	n := node.CreateNode(data)

	if s.firstNode == nil {
		s.firstNode = n
	} else {
		n.SetNext(s.firstNode)
		s.firstNode = n
	}
}

func (s *Stack) Pop() interface{} {
	if s.firstNode == nil {
		return nil
	}

	currNode := s.firstNode
	s.firstNode = s.firstNode.GetNext()

	return currNode.GetData()
}

func (s *Stack) IsEmpty() bool {
	return s.firstNode == nil
}
