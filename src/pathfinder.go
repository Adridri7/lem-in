package src

import "fmt"

var Paths [][]string

func GetAllValidPaths() {
	// How many from start (where can you go).
	for _, tunnel := range Anthill.Tunnels {
		if tunnel[0] == Anthill.Rooms[0].Index {
			Paths = append(Paths, tunnel)
		}
	}
	test := 0
	for test < 10 {
		for i, path := range Paths {
			count := 0
			// Init routes.
			for _, tunnel := range Anthill.Tunnels {
				if path[len(path)-1] == tunnel[0] && count < 1 {
					count++
					Paths[i] = append(Paths[i], tunnel...)
				} else if path[len(path)-1] == tunnel[0] {
					// Create a copy of the current path
					tempPath := make([]string, len(Paths[i])-2)
					copy(tempPath, Paths[i][:len(Paths[i])-2])
					// Append the tunnel to the copied path
					tempPath = append(tempPath, tunnel...)
					// Add the new path to Paths
					Paths = append(Paths, tempPath)
				}
			}
		}
		test++
	}
	fmt.Println(Paths)
}
