package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	escapeRoom, player := initialize()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {
		printPattern(&escapeRoom, &player)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyArrowUp { //  UP
				move(&escapeRoom, &player, North)
			} else if ev.Key == termbox.KeyArrowDown { // DOWN 
				move(&escapeRoom, &player, South)
			} else if ev.Key == termbox.KeyArrowLeft { // LEFT
				move(&escapeRoom, &player, West)
			} else if ev.Key == termbox.KeyArrowRight { // RIGTH 
				move(&escapeRoom, &player, East)
				if isGoal(escapeRoom, &player) {
					return
				}
			} else if ev.Ch == 'e' {  // 아이템 줍기 
				Interaction(&escapeRoom, &player)
			} else if ev.Ch == 'f' {  // 아이템 사용  
				itemEvent(&escapeRoom, &player)
			}else if ev.Key == termbox.KeyEsc {
				fmt.Println("ESC key pressed. Exiting...")
				termbox.Close()
			}
		}
	}
}
