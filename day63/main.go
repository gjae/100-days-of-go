package main

import "github.com/procyon-projects/chrono"
import "time"
import "context"
import "log"
import "os"

func main() {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	now := time.Now()

	_, err := taskScheduler.Schedule(func (ctx context.Context){
		f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
		_, _ = f.WriteString("One-shot task")
		_ = f.Close()
	}, chrono.WithStartTime(now.Year(), now.Month(),now.Day(), now.Hour(), now.Minute(), now.Second() + 1))
	
	if err == nil {
		log.Print("Task has been schedule successfully")
	}
	time.Sleep(time.Second * 2)
}
