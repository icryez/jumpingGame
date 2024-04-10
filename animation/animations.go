package animation

import (
	"fmt"
	colors "jumpingGame/colours"
	"jumpingGame/terminal"
	"jumpingGame/mapModule"
	"time"
)

var gameOver bool

func Tick() {
	gameOver = false
	for gameOver == false {
		time.Sleep(900 * time.Millisecond)
		mapModule.GenMap()
		grid:=mapModule.Grid
		terminal.CallClear()
		for r := range grid {
			for _, val := range grid[r] {
				if val.IsVisible {
					colors.Green.Print(" ")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}

}
