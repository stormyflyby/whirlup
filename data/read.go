package data

import (
	"fmt"
	"os"
)

func Read(dataPath string) string {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(data)
}
