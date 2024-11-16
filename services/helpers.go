package services

import (
	"fmt"
	"os"
)

func FileCreate(filePath string) {
	file, err := os.Create(filePath) // create a file
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}

	defer file.Close() // closing the file
}

// Read a file data and return it in a string format
func FileRead(filePath string) string {
	// Work here to read a json file data and return it to a string format
	//
	//
	//
	return ""
}
