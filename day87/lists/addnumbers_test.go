package lists

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	log.Print("Running tests")
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	var arrAssert []int
	result := []int{8, 9, 1}
	response := AddTwoNumbers(l1, l2)

	for response != nil {
		arrAssert = append(arrAssert, response.Val)
		response = response.Next
	}

	assert.ElementsMatch(t, result, arrAssert, "Arrays aren't equals")

}
