package stack

import (
	"fmt"
	"strconv"
)

type Stack struct {
	Val  string
	Next *Stack
}

func IsPalindrome(x int) bool {
	aux := fmt.Sprintf("%d", x)
	var root *Stack
	var result string

	for _, v := range aux {
		if root == nil {
			root = &Stack{Val: string(v)}
			continue
		}

		node := &Stack{Val: string(v), Next: root}
		root = node
	}

	for root != nil {
		result = fmt.Sprintf("%s%s", result, root.Val)
		root = root.Next
	}

	d, _ := strconv.Atoi(result)
	return d == x
}
