package src

import "fmt"

func Solve() {
	// Base case: if ending condition is met, return
	if AntsPop.EndRoomPop() == AntsPop.InitialPop {
		return
	}

	endingMoves := []string{}

	// Iterate over each ant's position
	for j, ant := range AntsPop.AntsPosition {
		moveCount := 0
		antMoved := false

		// Check possible paths for each ant
		for k := range LegitPaths {
			if k > 10 {
				break
			}
			for i, move := range LegitPaths[k] {
				if ant == move && i+1 < len(LegitPaths[k]) {
					nextRoom := LegitPaths[k][i+1]

					if !AntsPop.RoomIsFull(nextRoom) {
						if nextRoom == Anthill.EndRoom().Index {
							// Check if ant reaches end room
							if !containsString(endingMoves, ant) {
								endingMoves = append(endingMoves, ant)
								AntsPop.AntsPosition[j] = nextRoom
								moveCount++
								antMoved = true
								break
							}
						} else {

							if (len(LegitPaths[k])-i <= AntsPop.InitialPop-AntsPop.EndRoomPop()) || len(LegitPaths) < 2 {
								AntsPop.AntsPosition[j] = nextRoom
								moveCount++
								antMoved = true
							} else {
								if len(LegitPaths[0]) > AntsPop.InitialPop-AntsPop.EndRoomPop() {
									AntsPop.AntsPosition[j] = nextRoom
									moveCount++
									antMoved = true
								} else {
									antMoved = true
									break
								}
								AntsPop.AntsPosition[j] = nextRoom
								moveCount++
								antMoved = true
							}
						}
					}
				}
				if antMoved {
					break
				}
			}
			if antMoved {
				break
			}
		}

		// Print the movement if any
		if moveCount > 0 {
			fmt.Print("L", j+1, "-", AntsPop.AntsPosition[j], " ")
		}
	}

	fmt.Println()

	// Recursive call to continue solving
	Solve()
}

// Utility function to check if a string exists in a slice of strings
func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
