package services

import (
	"fmt"
	"time"

	"github.com/munnaMia/My-TODO-Manager/model"
)

// Add a task
func AddTask(title, description, taskFilePath string) {
	pendingTasksData := FileRead(taskFilePath) // All the pending task will be store here.

	taskSliceLength := len(pendingTasksData)

	newTask := model.Task{
		ID: taskSliceLength + 1,
		Title: title,
		Description: description,
		Completed: false,
		Created: time.Now(),
	}

	// Update the task in array 
	pendingTasksData = append(pendingTasksData, newTask)

	WriteFile(taskFilePath, pendingTasksData) // Update the file data.

	PrintSingleTask(pendingTasksData[taskSliceLength-1]) // Printing the test data
}

// Delete task by ID
func DeleteTaskByID(ID int) {
	fmt.Println("Deleted ID: ", ID)
}


func DeleteAllTask(filePath string) {
	RemoveAllData(filePath, "pending")
}


func DeleteCompletedTask(filePath string) {
	RemoveAllData(filePath, "completed")
}

// Completed task by ID
func CompletedTaskByID(id int) {
	fmt.Println("Completed task id:", id)
}

func ShowPendingTaskList(filePath string) {
	taskData := FileRead(filePath)
	PrintDATA(taskData)
}

func ShowCompletedTaskList(filePath string) {
	taskData := FileRead(filePath)
	PrintDATA(taskData)
}
