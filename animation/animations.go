package animation

import (
	"fmt"
	colors "jumpingGame/colours"
	"jumpingGame/mapModule"
	"jumpingGame/player"
	"jumpingGame/structs"
	"jumpingGame/terminal"
	"os"
	"os/exec"
	"time"
)

var GameInProgress bool
var FgridCurrCol int

func Tick() {
	GameInProgress = true
	mapModule.GenFgrid()
	terminal.CallClear()
	for GameInProgress {
		time.Sleep(10 * time.Millisecond)
		printGrid(mapModule.MainGrid)
	}

}

func printGrid(grid [30][100]structs.Cell) {

	// terminal.CallClear()

	terminal.MoveCursor(0, 0)
	for r := range grid {
		for c, val := range grid[r] {
			if player.PlayerCoordsMap.GetPlayCoord([2]int{r, c}, true) {
				colors.Yellow.Print(" ")
			} else if val.IsVisible {
				colors.Green.Print(" ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	terminal.CallFlush()
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
		time.Sleep(600 * time.Millisecond)
		for _, v := range player.PlayerCoordsMap.Keys() {
			if v[0] < 29 && v[0] >= 0 && !mapModule.MainGrid[v[0]][v[1]].IsVisible {
				player.PlayerCoordsMap.Lock()
				player.PlayerCoordsMap.SetPlayerCoord([2]int{v[0] + 2, v[1]}, true, false)
				player.PlayerCoordsMap.DeletePlayerCoords(v, false)
				player.PlayerCoordsMap.Unlock()
			} else {
				GameInProgress = false
			}
		}
	}
}

func ListenForJumps() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)

	for GameInProgress {
		os.Stdin.Read(b)
		if string(b) == " " {
			jump()
		}
	}
}

func jump() {
	for _, v := range player.PlayerCoordsMap.Keys() {
		if v[0] >= 0 && v[0] < 29 && !mapModule.MainGrid[v[0]][v[1]].IsVisible {
			player.PlayerCoordsMap.Lock()
			player.PlayerCoordsMap.SetPlayerCoord([2]int{v[0] - 1, v[1]}, true, false)
			player.PlayerCoordsMap.DeletePlayerCoords(v, false)
			player.PlayerCoordsMap.Unlock()
		} else {
			GameInProgress = false
		}
	}
}
