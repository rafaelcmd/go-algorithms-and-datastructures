package stacks

import "fmt"

//This solution uses a linked list structure to implement a stack

type Stack struct {
	Top *StackNode
}

type StackNode struct {
	Val  int
	Next *StackNode
}

func NewStack() *Stack {
	return &Stack{}
}

func NewStackNode(val int) *StackNode {
	return &StackNode{Val: val}
}

func (s *Stack) Push(val int) {
	node := NewStackNode(val)
	node.Next = s.Top
	s.Top = node
}

func (s *Stack) Pop() (int, error) {
	isEmpty := s.IsEmpty()
	if isEmpty {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.Top.Val
	s.Top = s.Top.Next
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	isEmpty := s.IsEmpty()
	if isEmpty {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.Top.Val, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Print() {
	current := s.Top
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
