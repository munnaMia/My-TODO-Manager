package services

import (
	"fmt"
	"time"

	"github.com/munnaMia/My-TODO-Manager/model"
)

// Add a task
func AddTask(title, description, taskFilePath string) {
	// Starting working on adding task.
	/*
		
		3. write to the storage
	*/
	// all pending task data 
	pendingTasksData := FileRead(taskFilePath) // All the pending task will be store here.

	newTask := model.Task{
		ID: len(pendingTasksData) + 1,
		Title: title,
		Description: description,
		Completed: false,
		Created: time.Now(),
	}

	// Update the task in array 
	pendingTasksData = append(pendingTasksData, newTask)

	for _, task := range pendingTasksData{
		fmt.Println(task)
	}
}

// Delete task by ID
func DeleteTaskByID(ID int) {
	fmt.Println("Deleted ID: ", ID)
}

// Delete all task
func DeleteAllTask() {
	fmt.Println("Deleted all task")
}

// Delete all completed task
func DeleteCompletedTask(filePath string) {
	RemoveAllData(filePath)
}

// Completed task by ID
func CompletedTaskByID(id int) {
	fmt.Println("Completed task id:", id)
}

// Show all pending task list
func ShowPendingTaskList() {
	fmt.Println("all task pending")
}

// Show all completed task list
func ShowCompletedTaskList() {
	fmt.Println("All completed pending last")
}
