// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func parseList(l *ListNode) []int {
	var response []int
	node := l

	for node != nil {
		response = append(response, node.Val)
		node = node.Next
	}

	return response
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := parseList(l1), parseList(l2)
	excedent := 0
	var newList *ListNode

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	var arrAux []int
	num2Len := len(num2)
	num1Len := len(num1)

	for i := 0; i <= num1Len-1; i++ {
		// aux := num1[i] + num2[i]
		aux := 0
		if num2Len > i {
			aux = num1[i] + num2[i]
		} else {
			aux = num1[i]
		}

		if excedent > 0 {
			aux += excedent
			excedent = 0
		}

		if aux < 10 {
			arrAux = append(arrAux, aux)
		} else if aux >= 10 && i < num1Len-1 {
			n := fmt.Sprintf("%d", aux)
			a, _ := strconv.Atoi(string(n[1]))
			b, _ := strconv.Atoi(string(n[0]))
			arrAux = append(arrAux, a)
			excedent += b
		} else {
			n := fmt.Sprintf("%d", aux)
			a, _ := strconv.Atoi(string(n[0]))
			b, _ := strconv.Atoi(string(n[1]))

			arrAux = append(arrAux, []int{b, a}...)
		}

	}

	for _, v := range arrAux {
		node := &ListNode{Val: v, Next: newList}
		newList = node
	}

	return newList
}

func main() {

	// l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	// l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}

	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}

	l3 := addTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}
