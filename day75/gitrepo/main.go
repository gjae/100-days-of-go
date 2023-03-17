package main

import (
	"example/commands/repo"
	"fmt"
)

func main() {

	command := repo.NewCommand("version")

	commandResult := command.Run("/home/giovanny/Projects/golang/100_days_of_go/", -1)
	fmt.Print(commandResult)
}
