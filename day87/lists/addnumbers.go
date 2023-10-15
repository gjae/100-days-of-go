package lists

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ParseList(l *ListNode) []int {
	var response []int
	node := l

	for node != nil {
		response = append(response, node.Val)
		node = node.Next
	}

	return response
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := ParseList(l1), ParseList(l2)
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
