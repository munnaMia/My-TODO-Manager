package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/munnaMia/My-TODO-Manager/services"
	"github.com/munnaMia/My-TODO-Manager/utils"
)

// Storage file paths...
var pendingTaskPath = filepath.Join("data", "task.json")
var completedTaskPath = filepath.Join("data", "completed.json")

func main() {
	utils.StorageFilesExist(pendingTaskPath)   // Checking storage files exists or not
	utils.StorageFilesExist(completedTaskPath) // Checking storage files exists or not

	if len(os.Args) < 2 {
		fmt.Println("Usages: todo <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("t", "", "Task of the Title")
		description := addCmd.String("d", "-", "Description of the Title")

		addCmd.Parse(os.Args[2:])

		if *title == "" {
			fmt.Println("Title is required. Usage: todo add -t <title> [-d <description>]")
			return
		}

		services.AddTask(*title, *description, pendingTaskPath)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		deleteByID := deleteCmd.Int("id", 0, "Delete a single task by ID")
		deleteAll := deleteCmd.Bool("all", false, "Delete all tasks")
		deleteCompleted := deleteCmd.Bool("c", false, "Delete all completed tasks")

		deleteCmd.Parse(os.Args[2:])

		// Checking how many flag are passing on delete cmd
		flagCounter := 0
		if *deleteAll {
			flagCounter++
		}
		if *deleteCompleted {
			flagCounter++
		}
		if *deleteByID != 0 {
			flagCounter++
		}
		if flagCounter > 1 {
			fmt.Println("Error: Only one of -all, -t <id>, or -c can be used at a time.")
			return
		}

		// Execute based on flags
		if *deleteAll {
			services.DeleteAllTask(pendingTaskPath)
		} else if *deleteCompleted {
			services.DeleteCompletedTask(completedTaskPath)
		} else if *deleteByID > 0 {
			services.DeleteTaskByID(*deleteByID)
		} else {
			fmt.Println("Please specify a flag: -t <id is greater then 0>, -all, or -c")
			return
		}

	case "complete":
		completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		completeID := completeCmd.Int("id", 0, "Completed task ID")

		if *completeID <= 0 {
			fmt.Println("Error: Task ID can't be a negative number or zero")
			return
		}

		services.CompletedTaskByID(*completeID)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		pendingTask := listCmd.Bool("p", false, "Showing all pending Task")
		completeTask := listCmd.Bool("c", false, "Showing all completed Task")

		listCmd.Parse(os.Args[2:])

		// Checking how many flag are passing on delete cmd
		flagCounter := 0
		if *pendingTask {
			flagCounter++
		}
		if *completeTask {
			flagCounter++
		}
		if flagCounter > 1 {
			fmt.Println("Error: Only one of -p, or -c can be used at a time.")
			return
		}

		// Execute based on flags
		if *pendingTask {
			services.ShowPendingTaskList(pendingTaskPath)
		} else if *completeTask {
			services.ShowCompletedTaskList(completedTaskPath)
		} else {
			fmt.Println("Please specify a flag: -p or -c")
			return
		}

	default:
		fmt.Println("Unknown command: ", command)
	}
}
