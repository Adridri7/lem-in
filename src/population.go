package src

type Population struct {
	InitialPop      int
	CurrentPop      int
	AntsCoordinates []Position
}

type Position struct {
	X, Y int
}

var AntsPop Population

// To check if an ant is already in a room.
func (p *Population) RoomIsFull(RoomIndex string) bool {
	for _, room := range Anthill.Rooms {
		if room.Index == RoomIndex {
			for _, antPos := range p.AntsCoordinates {
				if antPos.X == room.X && antPos.Y == room.Y {
					return true
				}
			}
		}
	}
	return false
}
