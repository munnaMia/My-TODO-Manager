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
	}
}
