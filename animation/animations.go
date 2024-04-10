package animation

import (
	"fmt"
	colors "jumpingGame/colours"
	"jumpingGame/mapModule"
	"jumpingGame/structs"
	"jumpingGame/terminal"
	"time"
)

var gameOver bool

func Tick() {
	gameOver = false
	for gameOver == false {
		time.Sleep(900 * time.Millisecond)
		mapModule.GenFgrid()
		printGrid(mapModule.Fgrid)
	}

}

func printGrid(grid [30][100]structs.Cell){

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
