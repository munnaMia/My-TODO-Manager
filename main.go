package main

import (
	"fmt"
	"os"
)

func main() {

	// If no command provided from user just return.
	if len(os.Args) < 2 {
		fmt.Println("Usages: todo <command> [arguments]")
		return
	}

	command := os.Args[1] // taking the command (add, list, delete...)

	switch command {
	case "add":
		fmt.Println("adding a task")
	case "delete":
		fmt.Println("deleting a task")
	case "complete":
		fmt.Println("completed a task")
	case "list":
		fmt.Println("printing tast pending/completed")
	default:
		fmt.Println("Unknown command: ", command)
	}
}
