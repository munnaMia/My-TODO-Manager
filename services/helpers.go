package services

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

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

	fileData, err := os.Open(filePath)
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

func RemoveAllData(filePath, typeOfTask string) {
	if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
		fmt.Printf("Failed to Delete Tasks : %v", err)
	}
	fmt.Printf("Successfully cleared all %v tasks.", typeOfTask)
}

func PrintDATA(tasks []model.Task) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(writer, "ID\tCREATED\tTITLE\tDESCRIPTION\t")
	fmt.Fprintln(writer, "----\t----\t----\t----\t")

	for _, task := range tasks {
		fmt.Fprintf(writer, "%v\t%v\t%v\t%v\t \n", task.ID, task.Created.Format("02-01-2006"), task.Title, task.Description)
	}

	writer.Flush()
}
