package services

import (
	"fmt"
)

// Add a task
func AddTask(title, description, taskFilePath string) {
	// Starting working on adding task.
	/*
		1. fetch all the JSON data from the storage.
		2. update the data
		3. write to the storage
	*/

	pendingTasks := FileRead(taskFilePath) // All the pending task will be store here.
	fmt.Println(pendingTasks)
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
func DeleteCompletedTask() {
	fmt.Println("Deleted all completed task")
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
