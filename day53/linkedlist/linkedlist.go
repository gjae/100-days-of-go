package main

import "fmt"

type Node struct {
	property int
	NextNode *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(firstProperty int) *LinkedList {
	node := &Node{property: firstProperty}
	linkedList := &LinkedList{}
	linkedList.Head = node

	return linkedList
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{property: property}
	node.NextNode = linkedList.Head
	linkedList.Head = node
}

func (linkedList *LinkedList) PopHead() *Node {
	node := linkedList.Head
	if node.NextNode != nil {
		linkedList.Head = node.NextNode
		return node
	}

	linkedList.Head = nil

	return nil
}

func (LinkedList LinkedList) LastNode() *Node {
	for node := LinkedList.Head; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			return node
		}
	}

	return nil
}

func (linkedList *LinkedList) AddToEnd(property int) {
	node := &Node{property: property}
	lastNode := linkedList.LastNode()

	if lastNode != nil {
		lastNode.NextNode = node
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWithValue *Node

	for node = linkedList.Head; node != nil; node = node.NextNode {
		if node.property == property {
			nodeWithValue = node
			break
		}
	}

	return nodeWithValue
}

func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{property: property}

	fNode := linkedList.NodeWithValue(nodeProperty)

	if fNode != nil {
		nodeAux := fNode.NextNode
		node.NextNode = nodeAux
		fNode.NextNode = node
	}
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.Head; node != nil; node = node.NextNode {
		fmt.Println(node.property)
	}

}

func main() {
	list := NewLinkedList(1)
	list.AddToHead(0)
	list.AddToEnd(3)
	list.AddAfter(1, 2)
	list.IterateList()
}
