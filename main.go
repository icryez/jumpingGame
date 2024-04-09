package main

import (
	"fmt"
	"jumpingGame/structs"
	"jumpingGame/colours"
	"jumpingGame/terminal"
	"math/rand"
)

var grid [30][100]structs.Cell
var nextObsStart int

func main() {
	terminal.CallClear()
	genMap()
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

func genMap() {
	for i := 20; i < 100; i++ {
		nextObsStart = i
		hTop := rand.Intn(20)
		genObs(hTop)
		i = i + 20
	}
}

func genObs(heigth int) {
	for k := 0; k <= heigth; k++ {
		for i := nextObsStart; i <= nextObsStart+6 && i < 100; i++ {
			grid[k][i].IsVisible = true
		}
	}
	for k:=heigth + 6;k<30;k++{
		for h:=nextObsStart;h<=nextObsStart+6 && h<100;h++{
			grid[k][h].IsVisible = true
		}
	}
}
