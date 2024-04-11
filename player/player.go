package player

import "sync"


var PlayerCoords map[[2]int]bool
var PlayerCoordsMutex sync.RWMutex
func PlayerStart() {
	createPlayerModel()
}

func createPlayerModel() {
	PlayerCoords = make(map[[2]int]bool)
	PlayerCoords[[2]int{14,6}] = true
	PlayerCoords[[2]int{15,6}] = true
	PlayerCoords[[2]int{14,7}] = true
	PlayerCoords[[2]int{15,7}] = true
	PlayerCoords[[2]int{14,8}] = true
	PlayerCoords[[2]int{15,8}] = true
}
