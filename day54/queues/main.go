package main

import "fmt"

type Order struct {
	Priority     int
	Quantity     int
	Product      string
	CustomerName string
}

type Queue []*Order

// NewOrder function initialize a new Order object
func NewOrder(priority, quantity int, product, customername string) *Order {
	return &Order{
		Priority:     priority,
		Quantity:     quantity,
		Product:      product,
		CustomerName: customername,
	}
}

// Add a new element to queue
func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		appened := false
		for i, addedOrder := range *q {
			if order.Priority > addedOrder.Priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appened = true
				break
			}
		}
		if !appened {
			*q = append(*q, order)
		}
	}
}

// main method
func main() {
	var queue Queue
	queue = make(Queue, 0)
	queue.Add(NewOrder(1, 2, "Orden: 00001", "Federico"))
	queue.Add(NewOrder(1, 2, "Orden: 00002", "Jaimito"))
	queue.Add(NewOrder(2, 2, "Orden: 00003", "Arianna"))

	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
