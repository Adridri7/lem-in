package src

import (
	"fmt"
	"os"
	"sort"
)

var LegitPaths [][]string

func GetAllValidPaths() {
	var Paths [][]string
	startRoom := ""
	endRoom := ""
	// Store start and end rooms.
	for _, room := range Anthill.Rooms {
		if room.isStartRoom {
			startRoom = room.Index
		}
		if room.isEndRoom {
			endRoom = room.Index
		}
	}
	// Copy + reverse tunnels to have bidirectionnal tunnels.
	tunnels := make(map[string][]string)

	for _, tunnel := range Anthill.Tunnels {

		if _, ok := tunnels[tunnel[0]]; !ok {
			tunnels[tunnel[0]] = []string{}
		}
		if _, ok := tunnels[tunnel[1]]; !ok {
			tunnels[tunnel[1]] = []string{}
		}
		tunnels[tunnel[0]] = append(tunnels[tunnel[0]], tunnel[1])
		tunnels[tunnel[1]] = append(tunnels[tunnel[1]], tunnel[0])
	}

	// Recursive function to find all possible paths.
	var findPaths func(currentRoom string, currentPath []string)
	findPaths = func(currentRoom string, currentPath []string) {
		currentPath = append(currentPath, currentRoom)
		if currentRoom == endRoom {
			temp := make([]string, len(currentPath))
			copy(temp, currentPath)
			Paths = append(Paths, temp)
		} else {
			for _, nextRoom := range tunnels[currentRoom] {
				if !contains(currentPath, nextRoom) { // Éviter les cycles
					findPaths(nextRoom, currentPath)
				}
			}
		}
	}
	// Start the research from startRoom
	findPaths(startRoom, []string{})
	// Handle error.
	if len(Paths) == 0 {
		fmt.Println("ERROR: No paths were found between start and end room")
		os.Exit(0)
	}
	// Sort paths by length.
	LegitPaths = append(LegitPaths, Paths...)
	SortByLength()
	fmt.Println(LegitPaths)
	PrintPaths()

}
func SortByLength() {
	sort.Slice(LegitPaths, func(i, j int) bool {
		return len(LegitPaths[i]) < len(LegitPaths[j])
	})
}
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func PrintPaths() {
	fmt.Println()
	for _, path := range LegitPaths {
		fmt.Println("Links:")
		for i := 0; i < len(path)-1; i++ {
			fmt.Printf("%s -> ", path[i])
		}
		fmt.Println(path[len(path)-1]) // Dernière salle sans flèche
	}
	fmt.Println()
}
