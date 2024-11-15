package services

import (
	"fmt"
	"os"
)

func FileCreate(filePath string) {
	// checking the task file existing or not
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath) // create a file
		if err != nil {
			fmt.Printf("Failed to create file: %v\n", err)
			return
		}

		defer file.Close() // closing the file
	} 
}
