package data

import (
	"log"
	"os"
	"path/filepath"

	"github.io/stormyflyby/whirlup/constants"
)

func Init() (dataPath string) {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		log.Fatalln(dirErr)
	}
	dataPath = filepath.Join(dir, constants.DataDirectory, constants.DataFile)
	mkDirErr := os.MkdirAll(filepath.Dir(dataPath), 0755)
	if mkDirErr != nil {
		log.Fatalf("Failed to make directory: %s\n", mkDirErr.Error())
	}
	fileErr := os.WriteFile(dataPath, []byte("{}"), 0755)
	if fileErr != nil {
		log.Fatalf("Failed to write to file: %s\n", fileErr.Error())
	}
	return dataPath
}
