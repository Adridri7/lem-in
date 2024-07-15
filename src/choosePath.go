/*package src

import "fmt"

// Count number of tunnels and ants, important because it's only one ant in a room

func CountTunnelsAndAnts(path []string, colony *Colony, antsPop *Population) int {
	count := 0
	for _, i := range path {
		if antsPop.RoomIsFull(i) {
			count++
		}
	}
	fmt.Println(count)
	return count
}

func ChoosePath(colony *Colony, antsPop *Population) []string {
	if len(Paths) < 2 {
		return Paths[0]
	}
	count1 := CountTunnelsAndAnts(Paths[0], colony, antsPop)
	count2 := CountTunnelsAndAnts(Paths[1], colony, antsPop)

	if count1 > count2 {
		return Paths[0]
	} else {
		return Paths[1]
	}
}

func SendAntInPath(colony *Colony, antsPop *Population) {
	GetAllValidPaths() // Populate Paths

	if len(Paths) == 0 {
		fmt.Println("Pas de chemin valide pour envoyer la fourmi")
		return
	}

	direction := ChoosePath(colony, antsPop)

	if len(direction) < 2 {
		fmt.Println("Pas de chemin valide pour envoyer la fourmi")
		return
	}

	nextDirection := direction[1]

	for i, pos := range antsPop.AntsCoordinates {
		// Trouvez la fourmi dans la pièce de départ
		if colony.Rooms[0].X == pos.X && colony.Rooms[0].Y == pos.Y {
			// Trouvez la pièce correspondante et mettez à jour les coordonnées
			for _, room := range colony.Rooms {
				if room.Index == nextDirection {
					antsPop.AntsCoordinates[i].X = room.X
					antsPop.AntsCoordinates[i].Y = room.Y
					fmt.Printf("Fourmi déplacée vers la pièce %s (%d, %d)\n", nextDirection, room.X, room.Y)
					return
				}
			}
		}
	}
}
*/

package src

import "fmt"

// DistributeAnts assigns ants to paths based on the given logic.
func DistributeAnts(numAnts int, paths [][]string) []int {
	pathAnts := make([]int, len(paths)) // to store the number of ants in each path
	for i := 0; i < numAnts; i++ {
		minPathIndex := 0
		minSteps := len(paths[0]) + pathAnts[0]
		for j := 1; j < len(paths); j++ {
			currentSteps := len(paths[j]) + pathAnts[j]
			if currentSteps < minSteps {
				minSteps = currentSteps
				minPathIndex = j
			}
		}
		pathAnts[minPathIndex]++
	}
	return pathAnts
}

// MoveAnts simulates the movement of ants through the paths.

func MoveAnts(paths [][]string, pathAnts []int) {
	// Track which rooms are occupied by ants
	occupiedRooms := make(map[string]bool)

	// Simulate the movement of ants
	for step := 0; ; step++ {
		allDone := true
		for pathIndex, numAnts := range pathAnts {
			for ant := 0; ant < numAnts; ant++ {
				if step < len(paths[pathIndex]) {

					roomIndex := paths[pathIndex][step]

					if step == 0 || !occupiedRooms[roomIndex] {
						fmt.Println("======================")
						fmt.Printf("Step %d: Ant %d on path %d in room %s\n", step, ant+1, pathIndex+1, roomIndex)
						occupiedRooms[roomIndex] = true
					} else {
						fmt.Printf("Step %d: Ant %d on path %d, room %s is full\n", step, ant+1, pathIndex+1, roomIndex)
					}
					allDone = false
				}
			}
		}
		if allDone {
			break
		}
	}
}

/*
package src

import "fmt"


func Queue(paths [][]string, antsAmount int) []int {

	antsQueue := make([]int, len(paths))
	rooms := make([]int, len(paths))

	for i := range paths {
		rooms[i] = len(paths[i])
	}

	for antsAmount > 0 {
		fmt.Println("rooms: ", rooms)
		fmt.Println("paths: ", len(paths))
		fmt.Println("antsQueue: ", antsQueue)
		fmt.Println("antsAmount: ", antsAmount)

		indexOfInsert := checkLowestPath(rooms, antsQueue)
		antsQueue[indexOfInsert] += 1
		antsAmount -= 1
	}
	return antsQueue
}


func checkLowestPath(rooms []int, antsQueue []int) int {
	lowestValue := 10000
	lowestInd := 0
	for indOfPath := range rooms {
		// summ of rooms and ants in one path
		sum := rooms[indOfPath] + antsQueue[indOfPath]
		if sum < lowestValue {
			lowestValue = sum
			lowestInd = indOfPath
		}
	}

	return lowestInd
}
*/
