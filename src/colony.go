package src

import "fmt"

type Colony struct {
	Rooms   []RoomInfos
	Tunnels [][]string
}
type RoomInfos struct {
	Index                  string
	X, Y                   int
	isStartRoom, isEndRoom bool
}

var Anthill Colony

// Find end room.
func (c *Colony) EndRoom() RoomInfos {
	for _, room := range c.Rooms {
		if room.isEndRoom {
			return room
		}
	}
	return RoomInfos{}
}

// Find start room.
func (c *Colony) StartRoom() RoomInfos {
	for _, room := range c.Rooms {
		if room.isStartRoom {
			return room
		}
	}
	return RoomInfos{}
}

// Find max x pos for rooms.
func (c *Colony) MaxXRoomPos() (maxX int) {
	maxX = 0
	for _, room := range c.Rooms {
		if room.X > maxX {
			maxX = room.X
		}
	}
	return
}

// Find max y pos for rooms.
func (c *Colony) MaxYRoomPos() (maxY int) {
	maxY = 0
	for _, room := range c.Rooms {
		if room.Y > maxY {
			maxY = room.Y
		}
	}
	return
}

// Display 2d map of the colony.
func (c *Colony) DisplayColony() {
	for y := 0; y <= c.MaxYRoomPos(); y++ {
		for x := 0; x <= c.MaxXRoomPos(); x++ {
			found := false
			for _, room := range c.Rooms {
				if room.Y == y && room.X == x {
					fmt.Print("[" + room.Index + "]")
					found = true
					break
				}
			}
			if !found {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
