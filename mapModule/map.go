package mapModule

import (
	"jumpingGame/structs"
	"math/rand"
)

var Fgrid [30][100]structs.Cell
var nextObsStart int
var MainGrid [30][100]structs.Cell

func GenFgrid() {
	Fgrid = [30][100]structs.Cell{}
	for i := 0; i < 100; i++ {
		nextObsStart = i
		hTop := rand.Intn(20)
		GenObs(hTop, 2)
		i = i + 35
	}
}

func GenObs(heigth int, grid int) {

	for k := 0; k <= heigth; k++ {
		for i := nextObsStart; i <= nextObsStart+5 && i < 100; i++ {
			if grid == 2 {
				Fgrid[k][i].IsVisible = true
			}
		}
	}
	for k := heigth + 8; k < 30; k++ {
		for h := nextObsStart; h <= nextObsStart+5 && h < 100; h++ {
			if grid == 2 {
				Fgrid[k][h].IsVisible = true
			}
		}
	}
}
