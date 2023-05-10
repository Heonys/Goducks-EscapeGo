package main

/*
docs
*/
type Room struct {
	door      [4]int
	item      string
	isMovable bool
	Required  string
	isEntry   bool
	isGoal    bool
	isBroken  bool
	pattern 	string
}


// 동 서 남 북
// 0은 못감
// 1은 바로 지나갈 수 있음
// 2는 문 열어야함
// 3은 망치 필요함
// 4는 열쇠 필요함

type Inven struct {
	hammer bool
	key    bool
}

type Playable struct {
	position [2]int
	inven    Inven
}

type Direction int

const (
	East Direction = iota
	West
	South
	North
)


