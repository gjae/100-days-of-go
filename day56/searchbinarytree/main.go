package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	key       int
	value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinarySearchTree struct {
	RootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	newNode := &TreeNode{key: key, value: value, LeftNode: nil, RightNode: nil}

	if tree.RootNode == nil {
		tree.RootNode = newNode
	} else {
		insertTree(tree.RootNode, newNode)
	}
}

func insertTree(rootNode *TreeNode, newNode *TreeNode) {
	if newNode.key < rootNode.key {
		if rootNode.LeftNode == nil {
			rootNode.LeftNode = newNode
		} else {
			insertTree(rootNode.LeftNode, newNode)
		}
	} else {
		if rootNode.RightNode == nil {
			rootNode.RightNode = newNode
		} else {
			insertTree(rootNode.RightNode, newNode)
		}
	}
}

func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	inOrderTraverseTree(tree.RootNode, function)
}

func inOrderTraverseTree(rootNode *TreeNode, function func(int)) {
	if rootNode != nil {
		inOrderTraverseTree(rootNode.LeftNode, function)
		function(rootNode.value)
		inOrderTraverseTree(rootNode.RightNode, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	preOrderTraverserTree(tree.RootNode, function)
}

func preOrderTraverserTree(rootNode *TreeNode, function func(int)) {
	if rootNode != nil {
		function(rootNode.value)
		preOrderTraverserTree(rootNode.LeftNode, function)
		preOrderTraverserTree(rootNode.RightNode, function)
	}
}

func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	postOrderTraverseTree(tree.RootNode, function)
}

func postOrderTraverseTree(rootNode *TreeNode, function func(int)) {
	if rootNode != nil {
		preOrderTraverserTree(rootNode.LeftNode, function)
		preOrderTraverserTree(rootNode.RightNode, function)
		function(rootNode.value)
	}
}

func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	var treeNode *TreeNode
	treeNode = tree.RootNode
	if treeNode == nil {
		return (*int)(nil)
	}

	for {
		if treeNode.LeftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.LeftNode
	}
}

func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.RootNode
	if treeNode == nil {
		return (*int)(nil)
	}

	for {
		if treeNode.RightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.RightNode
	}
}

func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	return searchNode(tree.RootNode, key)
}

func searchNode(node *TreeNode, key int) bool {

	if node == nil {
		return false
	}

	if node.key > key {
		return searchNode(node.LeftNode, key)
	} else if node.key < key {
		return searchNode(node.RightNode, key)
	}

	return true
}

func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.RootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode {

	if treeNode == nil {
		return nil
	}

	if treeNode.key > key {
		treeNode.LeftNode = removeNode(treeNode.LeftNode, key)
		return treeNode
	} else if treeNode.key < key {
		treeNode.RightNode = removeNode(treeNode.RightNode, key)
		return treeNode
	}

	if treeNode.LeftNode == nil && treeNode.RightNode == nil {
		treeNode = nil
		return nil
	}

	if treeNode.LeftNode == nil {
		treeNode = treeNode.RightNode
		return treeNode
	}

	if treeNode.RightNode == nil {
		treeNode = treeNode.LeftNode
		return treeNode
	}

	leftmostrightNode := treeNode.RightNode
	for {
		// find smallest value on the right side
		if leftmostrightNode != nil && leftmostrightNode.LeftNode != nil {
			leftmostrightNode = leftmostrightNode.LeftNode
		} else {
			break
		}
	}

	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.RightNode = removeNode(treeNode.RightNode, key)

	return treeNode
}

func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("------------------------------------")
	stringify(tree.RootNode, 0)
	fmt.Println("------------------------------------")
}

// stringify method
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---[ "
		level++

		// Print key inorder traversing
		stringify(treeNode.LeftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.RightNode, level)
	}
}

// main method
func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(10, 10)
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.InsertElement(0, 9)
	tree.String()
}
