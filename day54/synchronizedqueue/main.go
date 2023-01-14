package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants
const (
	MessagePassStart = iota
	MessageTicketStart
	MessagePassEnd
	MessageTicketEnd
)

// Queue struct
type Queue struct {
	WaitPass    int
	WaitTicket  int
	PlayPass    bool
	PlayTicket  bool
	QueuePass   chan int
	QueueTicket chan int
	Message     chan int
}

// New method initialize queue
func NewQueue() *Queue {
	queue := &Queue{
		Message:     make(chan int),
		QueuePass:   make(chan int),
		QueueTicket: make(chan int),
	}

	go func() {
		var message int
		for {
			select {
			case message = <-queue.Message:
				switch message {
				case MessagePassStart:
					queue.WaitPass++
				case MessagePassEnd:
					queue.PlayPass = false
				case MessageTicketStart:
					queue.WaitTicket++
				case MessageTicketEnd:
					queue.PlayTicket = false
				}
				if queue.WaitPass > 0 && queue.WaitTicket > 0 && !queue.PlayPass && !queue.PlayTicket {
					queue.PlayPass = true
					queue.PlayTicket = true
					queue.WaitTicket--
					queue.WaitPass--
					queue.QueuePass <- 1
					queue.QueueTicket <- 1
				}
			}
		}
	}()

	return queue
}

// StartTicketIssue starts the ticket issue
func (q *Queue) StartTicketIssue() {
	q.Message <- MessageTicketStart
	<-q.QueueTicket
}

func (q *Queue) EndTicketIssue() {
	q.Message <- MessageTicketEnd
}

func TicketIssue(q *Queue) {
	for {
		// Sleep up to 10 seconds
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		q.StartTicketIssue()
		fmt.Println("Ticket issue starts")

		// Sleep up to 2 seconds
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket issue ends")
		q.EndTicketIssue()
	}
}

func (q *Queue) StartPass() {
	q.Message <- MessagePassStart
	<-q.QueuePass
}

func (q *Queue) EndPass() {
	q.Message <- MessagePassEnd
}

// Passenger methods starts and ends the pass queue
func Passenger(q *Queue) {
	for {
		// Slep up to 10 seconds
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		q.StartPass()
		fmt.Println(" Passenger starts")

		// Sleep up to 2 seconds
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")
		q.EndPass()
	}
}

func main() {
	queue := NewQueue()
	for i := 0; i < 10; i++ {
		go Passenger(queue)
	}

	for j := 0; j < 5; j++ {
		go TicketIssue(queue)
	}

	select {}
}
