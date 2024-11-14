package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munnaMia/My-TODO-Manager/services"
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
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		titleFlag := addCmd.String("t", "", "Title of the task")
		descriptionFlag := addCmd.String("d", "No-Description", "Description of the task")

		addCmd.Parse(os.Args[2:])

		if *titleFlag == ""{
			fmt.Println("Title is required. Usage: todo add -t <title> [-d <description>]")
			return
		}

		services.AddTask(*titleFlag, *descriptionFlag)

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
