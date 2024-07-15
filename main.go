package main

import (
	"fmt"
	"os"
	"src"
)

func main() {
	args := os.Args[1:]
	src.Init(src.FileToTab(args[0]))
	fmt.Println(src.Anthill)
	src.Anthill.DisplayColony()
}
