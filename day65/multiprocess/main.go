package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"runtime"
)

func worker(id int) {
	data := []int{1,2,3,4,5,6,7,8,9,10, 11, 12, 13}
	fmt.Printf("Worker 0 PID: %d\n", os.Getpid())
	for _, v := range data {

                fmt.Printf("Worker %d: %d\n" , id, v,)
                time.Sleep(2 * time.Second)
        }
}


func main(){
	data := []int{1,2,3,4,5,6,7,8,9,10, 11, 12, 13}
	
	go worker(1)
	fmt.Printf("Worker 0 PID: %d\nWith Gorutines: %d\n", os.Getpid(), runtime.NumGoroutine())
	go func(){
        	dateCmd := exec.Command("./mainproc")
        	dateOut, _:= dateCmd.Output()

        	fmt.Println("> date ", string(dateOut))

	}()
	
	for _, v := range data {

		fmt.Println("Worker 0: ", v,)
		time.Sleep(2 * time.Second)
	}
}
