// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"gjae/linkedlist1/hash_maps"
	"gjae/linkedlist1/lists"
)

func main() {
	l1 := &lists.ListNode{Val: 9, Next: &lists.ListNode{Val: 9}}
	l2 := &lists.ListNode{Val: 9, Next: &lists.ListNode{Val: 9}}

	l3 := lists.AddTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

	fmt.Println(hash_maps.RunTwoNums([]int{2, 7, 11, 15}, 9))
	fmt.Println(hash_maps.RunTwoNums([]int{2, 4, 6, 7}, 8))
	fmt.Println(hash_maps.RunTwoNums([]int{3, 2, 4}, 6))

}
