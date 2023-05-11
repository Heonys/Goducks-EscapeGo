package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	escapeRoom, player := initalize()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()



	fmt.Println(tutotrial)
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
			if ev.Key == termbox.KeySpace {
				break
			} else if ev.Key == termbox.KeyEsc {
					fmt.Println("ESC key pressed. Exiting...")
					termbox.Close()
					return
			} 
	}


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
			} else if ev.Ch == 'e' { // 아이템 줍기 
				Interaction(&escapeRoom, &player)
			} else if ev.Ch == 'k' { 
				itemEvent(&escapeRoom, &player)
			}else if ev.Key == termbox.KeyEsc {
				fmt.Println("ESC key pressed. Exiting...")
				termbox.Close()
			}
		}
	}
}
