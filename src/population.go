package src

type Population struct {
	InitialPop   int
	CurrentPop   int
	AntsPosition []string
}

var AntsPop Population

// Init start room population.
func (p *Population) InitStartPop() {
	for i := 1; i <= p.InitialPop; i++ {
		p.AntsPosition = append(p.AntsPosition, Anthill.StartRoom().Index)
	}
}

// Check end room population.
func (p *Population) EndRoomPop() int {
	count := 0
	for _, antPos := range p.AntsPosition {
		if Anthill.EndRoom().Index == antPos {
			count++
		}
	}
	return count
}

// To check if an ant is already in a room.
func (p *Population) RoomIsFull(RoomIndex string) bool {
	if Anthill.EndRoom().Index == RoomIndex {
		return false
	}
	for _, pos := range p.AntsPosition {
		if pos == RoomIndex {
			return true
		}
	}
	return false
}
