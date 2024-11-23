package services

import (
	"fmt"
	"time"

	"github.com/munnaMia/My-TODO-Manager/model"
)

func AddTask(title, description, taskFilePath string) {
	pendingTasksData := FileRead(taskFilePath) // All the pending task will be store here.

	taskSliceLength := len(pendingTasksData)

	newTask := model.Task{
		ID:          taskSliceLength + 1,
		Title:       title,
		Description: description,
		Completed:   false,
		Created:     time.Now(),
	}

	// Update the task in array
	pendingTasksData = append(pendingTasksData, newTask)

	WriteFile(taskFilePath, "added", pendingTasksData) // Update the file data.

	PrintSingleTask(pendingTasksData[taskSliceLength]) // Printing the test data
}

// Delete task by ID
func DeleteTaskByID(ID int, filePath string) {
	pendingTasksData := FileRead(filePath)

	taskIdx, err := SearchTask(ID, pendingTasksData)
	if err != nil {
		fmt.Println(err)
		return
	}

	pendingTasksData = DeleteTask(taskIdx, pendingTasksData)
	pendingTasksData = SortIDs(pendingTasksData)

	WriteFile(filePath, "deleted", pendingTasksData)

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
