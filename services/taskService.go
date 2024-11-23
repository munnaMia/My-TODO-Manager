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

	WriteFile(taskFilePath, pendingTasksData) // Update the file data.

	PrintSingleTask(pendingTasksData[taskSliceLength]) // Printing the test data
	StatusPrint("added")
}

func DeleteTaskByID(ID int, filePath string) {
	pendingTasksData := FileRead(filePath)

	taskIdx, err := SearchTask(ID, pendingTasksData)
	if err != nil {
		fmt.Println(err)
		return
	}

	pendingTasksDatas, deletedTask := DeleteTask(taskIdx, pendingTasksData)
	pendingTasksData = SortIDs(pendingTasksDatas)

	PrintSingleTask(deletedTask)
	WriteFile(filePath, pendingTasksData)
	StatusPrint("deleted")

}

func DeleteAllTask(filePath string) {
	RemoveAllData(filePath, "pending")
}

func DeleteCompletedTask(filePath string) {
	RemoveAllData(filePath, "completed")
}

// Completed task by ID
func CompletedTaskByID(ID int, pendingTaskPath, completedTaskPath string) {
	pendingTaskFile := FileRead(pendingTaskPath)
	completeTasksFile := FileRead(completedTaskPath)

	// Searching the task exist or not
	taskIdx, err := SearchTask(ID, pendingTaskFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// task delete from pending task storage
	updatedPendingTasks, completedTask := DeleteTask(taskIdx, pendingTaskFile)
	pendingTaskFile = SortIDs(updatedPendingTasks)
	WriteFile(pendingTaskPath, pendingTaskFile)

	taskSliceLength := len(completeTasksFile)

	completeNewTask := model.Task{
		ID:          taskSliceLength + 1,
		Title:       completedTask.Title,
		Description: completedTask.Description,
		Completed:   true,
		Created:     completedTask.Created,
	}

	completeTasksFile = append(completeTasksFile, completeNewTask)
	WriteFile(completedTaskPath, completeTasksFile)
	PrintSingleTask(completedTask)
	StatusPrint("completed")
}

func ShowPendingTaskList(filePath string) {
	taskData := FileRead(filePath)
	PrintDATA(taskData)
}

func ShowCompletedTaskList(filePath string) {
	taskData := FileRead(filePath)
	PrintDATA(taskData)
}

func DisplayHelp() {
	helpMessage := `
Task Manager Command Help:

1. Add a Task
   Usage: todo add -t "title" [-d "description"]
   Description: Adds a new task.
      -t: Specifies the task title (required).
      -d: Specifies the task description (optional).

2. Delete a Task
   Usage: todo delete -id <id>
   Description: Deletes a task by its ID.

3. Delete All Tasks
   Usage: todo delete -all
   Description: Deletes all tasks.

4. Delete All Completed Tasks
   Usage: todo delete -c
   Description: Deletes all tasks marked as completed.

5. Complete a Task
   Usage: todo complete -id <id>
   Description: Marks a task as completed by its ID.

6. List Pending Tasks
   Usage: todo list -p
   Description: Lists all pending tasks.

7. List Completed Tasks
   Usage: todo list -c
   Description: Lists all completed tasks.

8. Help
   Usage: todo help
   Description: Displays this help message.
`
	fmt.Println(helpMessage)
}
