package main

import "fmt"

type Node struct {
	Value    int
	NextNode *Node
}

type Stack struct {
	Head    *Node
	Counter int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element *Node) {
	s.Counter++
	if s.Counter == 0 {
		s.Head = element
		return
	}

	element.NextNode = s.Head
	s.Head = element
}

func (s *Stack) Pop() *Node {
	var node *Node
	if s.Counter > 0 {
		node = s.Head
		s.Head = node.NextNode
		node.NextNode = nil
		s.Counter--
	}

	return node
}

func main() {
	stack := NewStack()
	stack.Push(&Node{Value: 1})
	stack.Push(&Node{Value: 2})
	stack.Push(&Node{Value: 3})
	stack.Push(&Node{Value: 0})

	for node := stack.Pop(); node != nil; node = stack.Pop() {
		fmt.Println(node.Value)
	}

	fmt.Printf("Current stack counter: %d\n", stack.Counter)
}
