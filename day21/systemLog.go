package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err !=  nil {
		log.Println(err)
	} else {
		log.SetOutput(sysLog)
		log.Print("Everything is fine")
		fmt.Print("Log entry ready")
	}
}