package main

import (
	"fmt"

	"github.io/stormyflyby/whirlup/data"
)

func main() {
	dataPath := data.Init()
	fmt.Println(data.Read(dataPath))
}
