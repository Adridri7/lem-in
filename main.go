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
	paths := src.GetAllValidPaths()

	// Number of ants (for example)
	numAnts := 4

	// Distribute ants among the paths
	pathAnts := src.DistributeAnts(numAnts, paths)

	// Move ants through the paths
	src.MoveAnts(paths, pathAnts)
}

/*
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
	paths := src.GetAllValidPaths()

	// Distribute ants among the paths
	numAnts := 4 // Nombre d'ants à distribuer
	pathAnts := src.Queue(paths, numAnts)

	// Move ants through the paths respecting the constraints
	fmt.Println("Distribution des fourmis par chemin :", pathAnts)
}
*/
