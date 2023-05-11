package main

import (
	"fmt"
)

var move = func(escapeRoom *([10][10]Room), player *Playable, direction Direction) {

	X := &player.position[0]
	Y := &player.position[1]
	CurrentRoom := &escapeRoom[player.position[0]][player.position[1]]

	switch direction {
	case East:
		if !escapeRoom[*X+1][*Y].isMovable {
			return
		}
		CheckDoorStatus(CurrentRoom, direction, player)
	case West:
		if !escapeRoom[*X-1][*Y].isMovable {
			return
		}
		CheckDoorStatus(CurrentRoom, direction, player)
	case South:
		if !escapeRoom[*X][*Y-1].isMovable {
			return
		}
		CheckDoorStatus(CurrentRoom, direction, player)
	case North:
		if !escapeRoom[*X][*Y+1].isMovable {
			return
		}
		CheckDoorStatus(CurrentRoom, direction, player)
	}
}

var Interaction = func(escapeRoom *([10][10]Room), player *Playable) {
	X := player.position[0]
	Y := player.position[1]
	CurrentRoom := escapeRoom[X][Y]

	switch CurrentRoom.item {
	case "hammer":
		if(!player.inven.hammer){
			player.inven.hammer = true
			fmt.Println("망치 획득")
		}
	case "key":
		if(!player.inven.key){
			player.inven.key = true
			fmt.Println("키 획득")
		}
	}
}


var CheckDoorStatus = func(CurrentRoom *Room, direction Direction, player *Playable) {
	switch CurrentRoom.door[direction] {
	case 0:
		return
	case 1, 2:
		CheckDirection(direction, player)
	case 3:
		if !player.inven.hammer {
			return
		}
		if CurrentRoom.isBroken {
			CheckDirection(direction, player)
		}
	case 4:
		if !player.inven.key {
			return
		}
		if CurrentRoom.isBroken {
			CheckDirection(direction, player)
		}
	}
}


var itemEvent = func(escapeRoom *([10][10]Room), player *Playable){
	CurrentRoom := &escapeRoom[player.position[0]][player.position[1]]
	var flag int
	for _, num := range CurrentRoom.door {
	 	if num > 1 {
			flag = num
			break
		}
	}
	switch flag {
	case 2:
		fmt.Println("문 열기")
	case 3:
		fmt.Println("망치 사용")
		CurrentRoom.isBroken = true
	case 4:
		fmt.Println("키 사용")
		CurrentRoom.isBroken = true
	}	
}




var CheckDirection = func(direction Direction, player *Playable) {
	switch direction {
	case 0:
		player.position[0]++
	case 1:
		player.position[0]--
	case 2:
		player.position[1]--
	case 3:
		player.position[1]++
	}
}

var isGoal = func(escapeRoom [10][10]Room, player *Playable) bool {
	X := player.position[0]
	Y := player.position[1]
	if escapeRoom[X][Y].isGoal {
		fmt.Print("게임 종료")
		return true
	}
	return false
}

var printPattern = func(escapeRoom *([10][10]Room), player *Playable){
	CurrentRoom := &escapeRoom[player.position[0]][player.position[1]]
	fmt.Println(CurrentRoom.pattern)
}