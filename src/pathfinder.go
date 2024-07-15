package src

import (
	"fmt"
	"strings"
)

var Paths [][]string

func GetAllValidPaths() [][]string {
	var newPath [][]string
	// Store the end room
	endRoom := ""
	for _, room := range Anthill.Rooms {
		if room.isEndRoom {
			endRoom = room.Index
		}
	}
	fmt.Println(endRoom)
	// Copying and inverting the tunnels
	var tunnels [][]string
	for _, tunnel := range Anthill.Tunnels {
		tunnels = append(tunnels, []string{tunnel[0], tunnel[1]})
		tunnels = append(tunnels, []string{tunnel[1], tunnel[0]})
	}
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
			for _, tunnel := range tunnels {
				if (path[len(path)-1] == tunnel[0] && count < 1) && (path[len(path)-1] != endRoom) && !strings.Contains(strings.Join(path, " "), tunnel[1]) {
					count++
					Paths[i] = append(Paths[i], tunnel...)
				} else if (path[len(path)-1] == tunnel[0]) && (path[len(path)-1] != endRoom) && !strings.Contains(strings.Join(path, " "), tunnel[1]) {
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
	for _, path := range Paths {
		if path[len(path)-1] == endRoom {
			newPath = append(newPath, path)
		}
	}
	fmt.Println(newPath)
	return newPath
}
