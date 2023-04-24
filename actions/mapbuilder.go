package lemin

import (
	"strconv"
	"strings"
)

type RoomNode struct {
	Name     string
	XCo      int
	YCo      int
	StartPos bool
	EndPos   bool
	Tunnels  []*RoomNode
}

func MapBuilder(fileData []string) (int, []RoomNode) {
	var numOfAnts int
	var rooms []string
	var links []string
	var startRoom string
	var endRoom string

	// find number of ants
	numOfAnts, _ = strconv.Atoi(fileData[0])

	// decode the strings and assign them to their relevant information tag
	for i := 1; i < len(fileData); i++ {
		runeArr := []rune(fileData[i])
		if fileData[i] == "##start" {
			startRoom = fileData[i+1]
		} else if fileData[i] == "##end" {
			endRoom = fileData[i+1]
		} else if runeArr[0] != '#' && runeArr[0] != 'L' {
			if !IsLink(runeArr) {
				rooms = append(rooms, fileData[i])
			} else {
				links = append(links, fileData[i])
			}
		}
	}

	// Create the nodes and fill out their data starting with names and coords
	nodes := make([]RoomNode, len(rooms))
	for i := 0; i < len(rooms); i++ {
		if startRoom == rooms[i] {
			nodes[i].StartPos = true
		} else if endRoom == rooms[i] {
			nodes[i].EndPos = true
		}
		roomDetails := strings.Split(rooms[i], " ")
		nodes[i].Name = roomDetails[0]
		nodes[i].XCo, _ = strconv.Atoi(roomDetails[1])
		nodes[i].YCo, _ = strconv.Atoi(roomDetails[2])
	}

	// add the links to the room nodes
	for j := 0; j < len(links); j++ {
		linksDetails := strings.Split(links[j], "-")
		if IsRoom(linksDetails[0], nodes) && IsRoom(linksDetails[1], nodes) {
			nodeNum1 := GetRoomNum(linksDetails[0], nodes)
			nodeNum2 := GetRoomNum(linksDetails[1], nodes)
			nodes[nodeNum1].Tunnels = append(nodes[nodeNum1].Tunnels, &nodes[nodeNum2])
			nodes[nodeNum2].Tunnels = append(nodes[nodeNum2].Tunnels, &nodes[nodeNum1])
		}
	}
	return numOfAnts, nodes
}

func IsRoom(room string, rooms []RoomNode) bool {
	for _, v := range rooms {
		if v.Name == room {
			return true
		}
	}
	return false
}

func GetRoomNum(room string, rooms []RoomNode) int {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Name == room {
			return i
		}
	}
	return -1
}

func IsLink(runeArr []rune) bool {
	for _, v := range runeArr {
		if v == '-' {
			return true
		}
	}
	return false
}
