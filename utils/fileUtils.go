package utils

import (
	"path/filepath"

	"github.com/munnaMia/My-TODO-Manager/services"
)

func StorageFilesExist() {
	pendingTaskPath := filepath.Join("data", "task.json")
	completedTaskPath := filepath.Join("data", "completed.json")

	services.FileCreate(pendingTaskPath)
	services.FileCreate(completedTaskPath)
}
