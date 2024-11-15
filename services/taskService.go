package services

import "fmt"

// Add a task
func AddTask(title, description string) {
	fmt.Println("done", title, description)
}

// Delete task by ID
func DeleteTaskByID(ID int){
	fmt.Println("Deleted ID: ",ID)
}

// Delete all task
func DeleteAllTask(){
	fmt.Println("Deleted all task")
}

// Delete all completed task
func DeleteCompletedTask(){
	fmt.Println("Deleted all completed task")
}

// Completed task by ID
func CompletedTaskByID(id int){
	fmt.Println("Completed task id:",id)
}