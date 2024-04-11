package main

import (
	"jumpingGame/animation"
	"jumpingGame/mapModule"
	"time"
)


func main() {
	go animation.Tick()
	animation.FgridCurrCol = 0
	for true {
		time.Sleep(200* time.Millisecond)
		animation.AnimateToLeft()
		if animation.FgridCurrCol == 100 {
		animation.FgridCurrCol = 0
			mapModule.GenFgrid()
		}
	}
}

