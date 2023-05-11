package main

var initialize = func() ([10][10]Room, Playable) {
	escapeRoom := initMap()
	player := initPlayer()
	return escapeRoom, player
}

//  Y
//
//  9  ⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛
//  8  ⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛
//  7  ⬛⬛⬛⬛⬛⬜⬜⬜🟨⬛
//  6  ⬛⬛⬛⬛⬜⬜⬛⬛⬛⬛
//  5  ⬛⬛⬛⬛⬜⬛⬛⬛⬛⬛
//  4  ⬛🔨⬜⬜⬜⬛⬛⬛⬛⬛
//  3  ⬛⬛⬜⬛⬜⬜⬛⬛⬛⬛
//  2  ⬛⬛⬜⬛⬛🔑⬛⬛⬛⬛
//  1  ⬛⬛⬜⬛⬛⬛⬛⬛⬛⬛
//  0  ⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛
//     0  1  2  3  4  5  6  7  8  9  X

var initMap = func() [10][10]Room {
	var _Map = [10][10]Room{}
	_Map[1][4] = Room{pattern: ptn4, door: [4]int{1, 0, 0, 0}, isMovable: true, item: "hammer"}
	_Map[2][1] = Room{pattern: ptn1, door: [4]int{0, 0, 0, 1}, isMovable: true, isEntry: true}
	_Map[2][2] = Room{pattern: ptn2, door: [4]int{0, 0, 1, 1}, isMovable: true}
	_Map[2][3] = Room{pattern: ptn2, door: [4]int{0, 0, 1, 1}, isMovable: true}
	_Map[2][4] = Room{pattern: ptn3, door: [4]int{3, 1, 1, 0}, isMovable: true}
	_Map[3][4] = Room{pattern: ptn5, door: [4]int{1, 1, 0, 0}, isMovable: true}
	_Map[4][3] = Room{pattern: ptn7, door: [4]int{2, 0, 0, 1}, isMovable: true, Required: "open"}
	_Map[4][4] = Room{pattern: ptn6, door: [4]int{0, 1, 1, 1}, isMovable: true}
	_Map[4][5] = Room{pattern: ptn2, door: [4]int{0, 0, 1, 1}, isMovable: true}
	_Map[4][6] = Room{pattern: ptn10, door: [4]int{1, 0, 1, 0}, isMovable: true}
	_Map[5][2] = Room{pattern: ptn9, door: [4]int{0, 0, 0, 1}, isMovable: true, item: "key"}
	_Map[5][3] = Room{pattern: ptn8, door: [4]int{0, 1, 1, 0}, isMovable: true, Required: "open"}
	_Map[5][6] = Room{pattern: ptn12, door: [4]int{0, 1, 0, 2}, isMovable: true, Required: "open"}
	_Map[5][7] = Room{pattern: ptn13, door: [4]int{1, 0, 1, 0}, isMovable: true, Required: "open"}
	_Map[6][7] = Room{pattern: ptn14, door: [4]int{1, 1, 0, 0}, isMovable: true}
	_Map[7][7] = Room{pattern: ptn15, door: [4]int{4, 1, 0, 0}, isMovable: true, Required: "key"}
	_Map[8][7] = Room{pattern: ptn16, door: [4]int{0, 1, 0, 0}, isMovable: true, isGoal: true}
	return _Map
}

var initPlayer = func() Playable {
	player := Playable{
		position: [2]int{2, 1},
		inven:    Inven{},
	}
	return player
}
