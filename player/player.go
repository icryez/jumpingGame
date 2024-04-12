package player

import (
	"sync"
)

type PlayerCoords struct {
	sync.Mutex
	m map[[2]int]bool
}

var PlayerCoordsMap PlayerCoords

func PlayerStart() {
	createPlayerModel()
}

func createPlayerModel() {
	PlayerCoordsMap = *NewPlayerCoordsMap()
	PlayerCoordsMap.SetPlayerCoord([2]int{14, 6}, true,true)
	PlayerCoordsMap.SetPlayerCoord([2]int{15, 6}, true,true)
	PlayerCoordsMap.SetPlayerCoord([2]int{14, 7}, true,true)
	PlayerCoordsMap.SetPlayerCoord([2]int{15, 7}, true,true)
	PlayerCoordsMap.SetPlayerCoord([2]int{14, 8}, true,true)
	PlayerCoordsMap.SetPlayerCoord([2]int{15, 8}, true,true)
	// PlayerCoords = make(map[[2]int]bool)
	// PlayerCoords[[2]int{14, 6}] = true
	// PlayerCoords[[2]int{15, 6}] = true
	// PlayerCoords[[2]int{14, 7}] = true
	// PlayerCoords[[2]int{15, 7}] = true
	// PlayerCoords[[2]int{14, 8}] = true
	// PlayerCoords[[2]int{15, 8}] = true
}

func NewPlayerCoordsMap() *PlayerCoords {
	pc := new(PlayerCoords)
	pc.m = map[[2]int]bool{}
	return pc
}

func (m *PlayerCoords) GetPlayCoord(key [2]int, lock bool) bool {
	if lock {
		m.Lock()
	}
	defer func() {
		if(lock){
			m.Unlock()
		}
	}()
	// defer m.Unlock()
	return m.m[key]
}

func (m *PlayerCoords) SetPlayerCoord(key [2]int, val bool, lock bool) {
	if lock {
		m.Lock()
	}
	defer func() {
		if(lock){
			m.Unlock()
		}
	}()
	m.m[key] = val
}

func (m *PlayerCoords) DeletePlayerCoords(key [2]int, lock bool) {
	if lock {
		m.Lock()
	}
	defer func() {
		if(lock){
			m.Unlock()
		}
	}()

	delete(m.m, key)
}

func (m *PlayerCoords) Keys() [][2]int {
	m.Lock()
	defer m.Unlock()

	keys := make([][2]int, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}
