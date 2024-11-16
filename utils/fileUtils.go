package utils

import (
	"os"

	"github.com/munnaMia/My-TODO-Manager/services"
)

// Checking file exist or not if not exist then create one
func StorageFilesExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		services.FileCreate(path)
	}
}
