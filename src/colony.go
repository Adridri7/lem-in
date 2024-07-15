package src

var GColony Colony

type Colony struct {
	Rooms   []RoomInfos
	Tunnels [][]string
}

type RoomInfos struct {
	Index                  string
	X, Y                   int
	isStartRoom, isEndRoom bool
}
