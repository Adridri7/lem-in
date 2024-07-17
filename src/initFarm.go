package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Init(args []string) {
	AntsPop.InitialPop, _ = strconv.Atoi(args[0])
	if AntsPop.InitialPop < 1 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(0)
	}
	for i := 1; i <= len(args)-1; i++ {
		if args[i] == "##start" {
			val := strings.Split(args[i+1], " ")
			index := val[0]
			x, _ := strconv.Atoi(val[1])
			y, _ := strconv.Atoi(val[2])
			roomInfos := RoomInfos{Index: index, X: x, Y: y, isStartRoom: true, isEndRoom: false}
			i++
			Anthill.Rooms = append(Anthill.Rooms, roomInfos)
		} else if args[i] == "##end" {
			val := strings.Split(args[i+1], " ")
			index := val[0]
			x, _ := strconv.Atoi(val[1])
			y, _ := strconv.Atoi(val[2])
			roomInfos := RoomInfos{Index: index, X: x, Y: y, isStartRoom: false, isEndRoom: true}
			Anthill.Rooms = append(Anthill.Rooms, roomInfos)
			i++
		} else if strings.HasPrefix(args[i], "#") {
			continue
		} else {
			if strings.Contains(args[i], "-") {
				Anthill.Tunnels = append(Anthill.Tunnels, strings.Split(args[i], "-"))
			} else {
				val := strings.Split(args[i], " ")
				index := val[0]
				x, _ := strconv.Atoi(val[1])
				y, _ := strconv.Atoi(val[2])
				roomInfos := RoomInfos{Index: index, X: x, Y: y, isStartRoom: false, isEndRoom: false}
				Anthill.Rooms = append(Anthill.Rooms, roomInfos)
			}
		}
	}
	AntsPop.InitStartPop()
	fmt.Println(AntsPop.InitialPop)
}
