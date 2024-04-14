package main

import (
	"jumpingGame/animation"
	colors "jumpingGame/colours"
	"jumpingGame/mapModule"
	"jumpingGame/player"
	"jumpingGame/terminal"
	"os"
	"time"

	tm "laptudirm.com/x/terminal"
)

func main() {
	term := tm.New(os.Stdout)
	term.HideCursor()
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
	time.Sleep(200 * time.Millisecond)
	terminal.CallClearCmd()
	terminal.MoveCursor(0, 0)
	term.ShowCursor()
	colors.Red.Println("############# GAME OVER ############")
}
