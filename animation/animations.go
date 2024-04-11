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
var FgridCurrCol int
func Tick() {
	gameOver = false
	FgridCurrCol = 0
	mapModule.GenFgrid()
	for gameOver == false {
		
		time.Sleep(150 * time.Millisecond)
		animateToLeft()
		printGrid(mapModule.MainGrid)
		if FgridCurrCol == 100 {
			FgridCurrCol = 0
			mapModule.GenFgrid()
		}
	}

}

func printGrid(grid [30][100]structs.Cell) {

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

func animateToLeft() {
	for i := 0; i < 30; i++ {
		for j := 0; j < 99; j++ {
			mapModule.MainGrid[i][j] = mapModule.MainGrid[i][j+1]
		}
	}
	for i:=0;i<30;i++{
		mapModule.MainGrid[i][99]=mapModule.Fgrid[i][FgridCurrCol]
	}
	FgridCurrCol++
}
