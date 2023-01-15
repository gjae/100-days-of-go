package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	Value int
}

type Stack struct {
	elements     []*Element
	ElementCount int
}

func (e *Element) String() string {
	return strconv.Itoa(e.Value)
}

func (s *Stack) New() {
	s.elements = make([]*Element, 0)
}

func (s *Stack) Push(element *Element) {
	s.elements = append(s.elements[:s.ElementCount], element)
	s.ElementCount++
}

func (s *Stack) Pop() *Element {
	var element *Element

	if s.ElementCount > 0 {
		length := s.ElementCount
		element = s.elements[length-1]
		s.elements = s.elements[0 : length-1]
		s.ElementCount--
	}

	return element
}

func main() {
	stack := &Stack{}
	stack.New()
	stack.Push(&Element{1})
	stack.Push(&Element{2})
	stack.Push(&Element{3})
	stack.Push(&Element{0})

	for node := stack.Pop(); node != nil; node = stack.Pop() {
		if node != nil {
			fmt.Println(node.Value)
		}
	}

	fmt.Printf("Current stack counter: %d\n", stack.ElementCount)
}
