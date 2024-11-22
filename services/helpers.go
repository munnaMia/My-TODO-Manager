package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/munnaMia/My-TODO-Manager/model"
)

func FileCreate(filePath string) {
	file, err := os.Create(filePath) 
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}

	defer file.Close() 
}

// Read a file data and return it in a slice of map format
func FileRead(filePath string) []model.Task {
	
	var tasks []model.Task

	fileData , err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer fileData.Close()

	decoder := json.NewDecoder(fileData)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Printf("Error decoding JSON : %v", err)
		return nil
	}
	
	return tasks
}

func RemoveAllData(filePath string){
	if err := os.WriteFile(filePath, []byte(""), 0644); err != nil{
		fmt.Printf("Failed to Delete Tasks : %v", err)
	}
	fmt.Println("Successfully cleared tasks.")
}
