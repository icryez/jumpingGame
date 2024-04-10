package mapModule

import (
	"jumpingGame/structs"
	"math/rand"
)

var Fgrid [30][100]structs.Cell
var nextObsStart int

// var Sgrid [30][100]structs.Cell
var MainGrid [30][100]structs.Cell

// func GenSgrid() {
// 	Sgrid = [30][100]structs.Cell{}
// 	for i := 20; i < 100; i++ {
// 		nextObsStart = i
// 		hTop := rand.Intn(20)
// 		GenObs(hTop, 1)
// 		i = i + 20
// 	}
// }

func GenFgrid() {
	Fgrid = [30][100]structs.Cell{}
	for i := 20; i < 100; i++ {
		nextObsStart = i
		hTop := rand.Intn(20)
		GenObs(hTop, 2)
		i = i + 20
	}
}

func GenObs(heigth int, grid int) {

	for k := 0; k <= heigth; k++ {
		for i := nextObsStart; i <= nextObsStart+5 && i < 100; i++ {
			if grid == 2 {
				Fgrid[k][i].IsVisible = true
			}
			// } else {
			// 	Sgrid[k][i].IsVisible = true
			// }
		}
	}
	for k := heigth + 6; k < 30; k++ {
		for h := nextObsStart; h <= nextObsStart+5 && h < 100; h++ {
			if grid == 2 {
				Fgrid[k][h].IsVisible = true
			}
			// } else {
			// 	Sgrid[k][h].IsVisible = true
			// }
		}
	}
}
