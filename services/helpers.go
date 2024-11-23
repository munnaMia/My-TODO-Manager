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
	if len(tasks) == 0 {
		return
	}
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(writer, "ID\tCREATED\tTITLE\tDESCRIPTION\t")
	fmt.Fprintln(writer, "----\t----\t----\t----\t")

	for _, task := range tasks {
		fmt.Fprintf(writer, "%v\t%v\t%v\t%v\t \n", task.ID, task.Created.Format("02-01-2006"), task.Title, task.Description)
	}

	writer.Flush()
}

func PrintSingleTask(task model.Task) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(writer, "ID\tCREATED\tTITLE\tDESCRIPTION\t")
	fmt.Fprintln(writer, "----\t----\t----\t----\t")

	fmt.Fprintf(writer, "%v\t%v\t%v\t%v\t \n", task.ID, task.Created.Format("02-01-2006"), task.Title, task.Description)

	writer.Flush()
}

func WriteFile(filePath string, tasks []model.Task) {
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Error opening file for writing:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func StorageFilesExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		FileCreate(path)
	}
}

func SearchTask(ID int, tasks []model.Task) (int, error) {
	for idx, task := range tasks {
		if ID == task.ID {
			return int(idx), nil
		}
	}
	return 0, fmt.Errorf("this id %v doesn't exist", ID)
}

func DeleteTask(index int, tasks []model.Task) ([]model.Task, model.Task) {
	selectedTask := tasks[index]
	tasks = append(tasks[:index], tasks[index+1:]...)

	return tasks, selectedTask
}

func SortIDs(tasks []model.Task) []model.Task {
	for idx := range tasks {
		tasks[idx].ID = idx + 1
	}
	return tasks
}


func StatusPrint(status string){
	fmt.Printf("Task Successfully %v", status)
}