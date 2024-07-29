package pkg

import (
	"os"
)

// CheckIfFilesExist проверяет, есть ли файлы в указанной папке
func CheckIfFilesExist(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			return true, nil
		}
	}
	return false, nil
}
