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
			if player.PlayerCoordsMap.GetPlayCoord([2]int{r, c},true) {
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
		time.Sleep(500 * time.Millisecond)
		for _, v := range player.PlayerCoordsMap.Keys() {
			if v[0] < 29 {
				player.PlayerCoordsMap.Lock()
				player.PlayerCoordsMap.SetPlayerCoord([2]int{v[0] + 4, v[1]}, true, false)
				player.PlayerCoordsMap.DeletePlayerCoords(v,false)
				player.PlayerCoordsMap.Unlock()
			} else {
				GameInProgress = false
			}
		}
	}
}

func ListenForJumps() {
	for GameInProgress {
			jump()
	}
}

func jump() {
	time.Sleep(200 * time.Millisecond)
	for _, v := range player.PlayerCoordsMap.Keys() {
		if v[0] >= 0 {
			player.PlayerCoordsMap.Lock()
			player.PlayerCoordsMap.SetPlayerCoord([2]int{v[0] - 2, v[1]}, true, false)
			player.PlayerCoordsMap.DeletePlayerCoords(v,false)
			player.PlayerCoordsMap.Unlock()
		} else {
			GameInProgress = false
		}
	}
}
