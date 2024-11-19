package services

import (
	"encoding/json"
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

// Read a file data and return it in a slice of map format
func FileRead(filePath string) []map[string]interface{} {
	
	var tasks []map[string]interface{}

	fileData , err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	decoder := json.NewDecoder(fileData)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Println("Error decoding JSON : %v", err)
		return nil
	}
	
	return tasks
}
