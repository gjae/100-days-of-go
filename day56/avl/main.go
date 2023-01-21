package main

import (
	"encoding/json"
	"fmt"
)

// KeyValue Type
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualsTo(KeyValue) bool
}

// TreeNode type
type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

type integerKey int

func (k integerKey) LessThan(k1 KeyValue) bool { return k < k1.(integerKey) }
func (k integerKey) EqualsTo(k1 KeyValue) bool { return k == k1.(integerKey) }

// Opposite method
func Opposite(nodeValue int) int {
	return 1 - nodeValue
}

// Single rotation method
func SingleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[Opposite(nodeValue)]
	rootNode.LinkedNodes[Opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// Double rotation
func DoubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[Opposite(nodeValue)].LinkedNodes[nodeValue]
	rootNode.LinkedNodes[Opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[Opposite(nodeValue)]
	saveNode.LinkedNodes[Opposite(nodeValue)] = rootNode.LinkedNodes[Opposite(nodeValue)]
	rootNode.LinkedNodes[Opposite(nodeValue)] = saveNode
	saveNode = rootNode.LinkedNodes[Opposite(nodeValue)]
	rootNode.LinkedNodes[Opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// Adjust balance method
func AdjustBalance(rootNode *TreeNode, nodeVchan int, balanceValue int) {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeVchan]
	var appNode *TreeNode
	appNode = node.LinkedNodes[Opposite(balanceValue)]
	switch appNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}

	appNode.BalanceValue = 0
}

// BalanceTree method
func BalanceTree(rootNode *TreeNode, nodeValue int) *TreeNode {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return SingleRotation(rootNode, Opposite(nodeValue))
	}

	AdjustBalance(rootNode, nodeValue, balance)
	return DoubleRotation(rootNode, Opposite(nodeValue))
}

// InsertRNode method
func InsertRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}

	dir := 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = InsertRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, done
	}

	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}

	return BalanceTree(rootNode, dir), true
}

// InsertNode method
func InsertNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = InsertRNode(*treeNode, key)
}

// RemoveNode method
func RemoveNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = removeRNode(*treeNode, key)
}

// RemoveBalance method
func removeBalance(rootNode *TreeNode, nodeValue int) (*TreeNode, bool) {
	var node *TreeNode
	node = rootNode.LinkedNodes[Opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return SingleRotation(rootNode, nodeValue), false
	case balance:
		AdjustBalance(rootNode, Opposite(nodeValue), -balance)
		return DoubleRotation(rootNode, nodeValue), false
	}

	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return SingleRotation(rootNode, nodeValue), true
}

// removeRNode method
func removeRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}

	if rootNode.KeyValue.EqualsTo(key) {
		switch {
		case rootNode.LinkedNodes[0] == nil:
			return rootNode.LinkedNodes[1], false
		case rootNode.LinkedNodes[1] == nil:
			return rootNode.LinkedNodes[0], false
		}
		var heirNode *TreeNode
		heirNode = rootNode.LinkedNodes[0]
		for heirNode.LinkedNodes[1] != nil {
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = removeRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}

	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}

	return removeBalance(rootNode, dir)

}

//main method
func main() {
	var treeNode *TreeNode
	fmt.Println("Tree is empty")
	var avlTree []byte
	avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	fmt.Println(string(avlTree))

	fmt.Println("\n Add Tree")
	InsertNode(&treeNode, integerKey(5))
	InsertNode(&treeNode, integerKey(3))
	InsertNode(&treeNode, integerKey(8))
	InsertNode(&treeNode, integerKey(7))
	InsertNode(&treeNode, integerKey(6))
	InsertNode(&treeNode, integerKey(10))
	avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	fmt.Println(string(avlTree))

	fmt.Println("\n Delete Tree")
	RemoveNode(&treeNode, integerKey(3))
	RemoveNode(&treeNode, integerKey(7))
	avlTree, _ = json.MarshalIndent(treeNode, "", " ")
	fmt.Println(string(avlTree))
}
