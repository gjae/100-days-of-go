package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var integers []int = []int{1, 2, 3, 4, 5, 6}

	var circularList *ring.Ring
	circularList = ring.New(len(integers))

	var i int
	for i = 0; i < circularList.Len(); i++ {
		circularList.Value = integers[i]
		circularList = circularList.Next()
	}

	circularList.Do(func(el interface{}) {
		fmt.Print(el, " -> ")
	})

	fmt.Println()

	// Reverse circular list
	for i = 0; i < circularList.Len(); i++ {
		fmt.Print(circularList.Value, " -> ")
		circularList = circularList.Prev()
	}

	fmt.Println()

	circularList = circularList.Move(2)
	circularList.Do(func(el interface{}) {
		fmt.Print(el, " -> ")
	})
	fmt.Println()

}
