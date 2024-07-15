package src

import (
	"fmt"
	"strconv"
	"strings"
)

func Init(args []string) {
	antsNum := args[0]
	fmt.Println(args)
	end := false

	for i := 1; i <= len(args)-1; i++ {
		if args[i] == "##start" {
			val := strings.Split(args[i+1], " ")
			index := val[0]
			x, _ := strconv.Atoi(val[1])
			y, _ := strconv.Atoi(val[2])
			roomInfos := RoomInfos{Index: index, X: x, Y: y, isStartRoom: true, isEndRoom: false}
			i++
			GColony.Rooms = append(GColony.Rooms, roomInfos)
			fmt.Println(roomInfos)

		} else if args[i] == "##end" {
			val := strings.Split(args[i+1], " ")
			index := val[0]
			x, _ := strconv.Atoi(val[1])
			y, _ := strconv.Atoi(val[2])
			roomInfos := RoomInfos{Index: index, X: x, Y: y, isStartRoom: false, isEndRoom: true}
			GColony.Rooms = append(GColony.Rooms, roomInfos)
			i++
			end = true
			fmt.Println(roomInfos)
		} else {
			if end {
				GColony.Tunnels = append(GColony.Tunnels, strings.Split(args[i], "-"))
			} else {
				val := strings.Split(args[i], " ")
				index := val[0]
				x, _ := strconv.Atoi(val[1])
				y, _ := strconv.Atoi(val[2])
				roomInfos := RoomInfos{Index: index, X: x, Y: y, isStartRoom: false, isEndRoom: false}
				GColony.Rooms = append(GColony.Rooms, roomInfos)
				fmt.Println(roomInfos)
			}
		}
	}
	fmt.Println(antsNum)
}
