package data

import (
	"log"
	"os"
	"path/filepath"

	"github.io/stormyflyby/whirlup/constants"
)

func Init() (dirPath string) {
	configDir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		log.Fatalln(dirErr)
	}
	dirPath = filepath.Join(configDir, constants.DataDirectory)
	mkDirErr := os.MkdirAll(dirPath, 0755)
	if mkDirErr != nil {
		log.Fatalf("Failed to make directory: %s\n", mkDirErr.Error())
	}
	return dirPath
}
