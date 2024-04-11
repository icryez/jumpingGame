package animation

import (
	"fmt"
	colors "jumpingGame/colours"
	"jumpingGame/mapModule"
	"jumpingGame/player"
	"jumpingGame/structs"
	"jumpingGame/terminal"
	"time"
)

var GameInProgress bool
var FgridCurrCol int

func Tick() {
	GameInProgress = true
	mapModule.GenFgrid()
	for GameInProgress {
		time.Sleep(10 * time.Millisecond)
		printGrid(mapModule.MainGrid)
	}

}

func printGrid(grid [30][100]structs.Cell) {

	terminal.CallClear()
	for r := range grid {
		for c, val := range grid[r] {
			if player.PlayerCoords[[2]int{r, c}] {
				colors.Yellow.Print(" ")
			} else if val.IsVisible {
				colors.Green.Print(" ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func AnimateToLeft() {
	for i := 0; i < 30; i++ {
		for j := 0; j < 99; j++ {
			mapModule.MainGrid[i][j].IsVisible = mapModule.MainGrid[i][j+1].IsVisible
		}
	}
	for i := 0; i < 30; i++ {
		mapModule.MainGrid[i][99].IsVisible = mapModule.Fgrid[i][FgridCurrCol].IsVisible
	}
	FgridCurrCol++
}

func StartPlayerGravity() {
	for GameInProgress {
		for k := range player.PlayerCoords {
			time.Sleep(100 * time.Millisecond)
			if k[0] < 29 {
				player.PlayerCoords[[2]int{k[0] + 1, k[1]}] = true
				delete(player.PlayerCoords, k)
			} else {
				GameInProgress = false
			}
		}
	}
}
