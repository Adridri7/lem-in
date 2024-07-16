package main

import (
	"os"
	"src"
)

func main() {
	args := os.Args[1:]
	src.Init(src.FileToTab(args[0]))
	src.GetAllValidPaths()
	src.Anthill.DisplayColony()
	src.Solve()
}
