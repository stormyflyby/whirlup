package main

import (
	"fmt"

	"github.io/stormyflyby/whirlup/data"
)

func main() {
	dirPath := data.Init()
	filePath, err := data.CreateFile(dirPath, "test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data.Read(filePath))
}
