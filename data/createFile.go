package data

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(dirPath, fileName string) (filePath string, err error) {
	filePath = filepath.Join(dirPath, fileName)

	// check if file already exists
	_, err = os.Stat(filePath)
	if !errors.Is(err, os.ErrNotExist) {
		errMsg := fmt.Sprintf("file %s already exists", filePath)
		return "", errors.New(errMsg)
	}

	fileErr := os.WriteFile(filePath, []byte("{}\n"), 0755)
	if fileErr != nil {
		log.Fatalf("Failed to write to file: %s\n", fileErr.Error())
	}

	return filePath, nil
}
