package main

import "fmt"

type Node struct {
	Property int
	NextNode *Node
	PrevNode *Node
}

type LinkedList struct {
	Head *Node

	// Point to last node in the list
	Tail *Node
}

func (l *LinkedList) First() *Node {
	return l.Head
}

func (l *LinkedList) Last() *Node {
	return l.Tail
}

func (linkedList *LinkedList) Prepend(property int) {
	node := &Node{Property: property}

	if linkedList.Head == nil {
		linkedList.Head = node
		linkedList.Tail = node
		return
	}

	node.NextNode = linkedList.Head
	linkedList.Head.PrevNode = node
	linkedList.Head = node
}

func (l *LinkedList) Append(property int) {
	node := &Node{Property: property}
	lastNode := l.Last()

	node.PrevNode = lastNode
	lastNode.NextNode = node

	l.Tail = node
}

func (l *LinkedList) Find(property int) *Node {
	var node *Node
	var nodeValue *Node

	for node = l.Head; node != nil; node = node.NextNode {
		if node.Property == property {
			nodeValue = node
		}
	}

	return nodeValue
}

func (l *LinkedList) InsertAfter(after, insert int) {
	node := l.Find(after)
	newNode := &Node{Property: insert}

	if node == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	if node.NextNode == nil {
		node.NextNode = newNode
		newNode.PrevNode = node
		l.Tail = newNode
		return
	}

	newNode.NextNode = node.NextNode

	node.NextNode.PrevNode = newNode

	newNode.PrevNode = node
	node.NextNode = newNode
}

func (l *LinkedList) IterateList() {
	for node := l.Head; node != nil; node = node.NextNode {
		fmt.Println(node.Property)
	}
}

func (l *LinkedList) IterateReverse() {
	for node := l.Last(); node != nil; node = node.PrevNode {
		fmt.Println(node.Property)
	}
}

func main() {
	linkedList := &LinkedList{}
	linkedList.Prepend(1)
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.InsertAfter(2, 3)
	linkedList.InsertAfter(4, 5)
	linkedList.Prepend(0)

	linkedList.IterateList()
	fmt.Println("Iterate in reverse mode ")
	linkedList.IterateReverse()
}
