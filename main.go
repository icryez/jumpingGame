package main

import (
	"jumpingGame/animation"
	"jumpingGame/mapModule"
	"jumpingGame/player"
	"time"
)

func main() {
	animation.GameInProgress = true
	go animation.Tick()
	player.PlayerStart()
	animation.FgridCurrCol = 0
	go animation.StartPlayerGravity()
	for animation.GameInProgress {
		time.Sleep(200* time.Millisecond)
		animation.AnimateToLeft()
		if animation.FgridCurrCol == 100 {
		animation.FgridCurrCol = 0
			mapModule.GenFgrid()
		}
	}
}

