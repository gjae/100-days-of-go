package main

import (
	"fmt"
	"os"
	"strconv"
	"context"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("F1() Done: ", c1.Err())
		return
	case r := <- time.After(time.Duration(t) * time.Second):
		fmt.Println("f1(): ", r)
	}
	
	return
}


func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2() Done: ", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2(): ", r)
	}

	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)

	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3() Done: ", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3(): ", r)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give a duration as int type ")
		return
	}

	t, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return 
	}

	f1(t)
	f2(t)
	f3(t)
}