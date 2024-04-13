package main

import (
	"jumpingGame/animation"
	colors "jumpingGame/colours"
	"jumpingGame/mapModule"
	"jumpingGame/player"
	"jumpingGame/terminal"
	"time"
)

func main() {
	animation.GameInProgress = true
	go animation.Tick()
	player.PlayerStart()
	animation.FgridCurrCol = 0
	go animation.StartPlayerGravity()
	go animation.ListenForJumps()
	for animation.GameInProgress {
		time.Sleep(200 * time.Millisecond)
		animation.AnimateToLeft()
		if animation.FgridCurrCol == 100 {
			animation.FgridCurrCol = 0
			mapModule.GenFgrid()
		}
	}
	time.Sleep(100 * time.Millisecond)
	terminal.CallClear()
	colors.Red.Println("############# GAME OVER ############")
}
