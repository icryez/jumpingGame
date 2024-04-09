package main

import (
	"fmt"
	"jumpingGame/structs"
)

var grid [30][100]structs.Cell

func main(){
	for r:= range grid {
		for _,val := range grid[r] {
			if val.IsVisible == false{
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func genMap(){
	for r:= range grid {
		for c,val := range grid[r]{
			genRandomNumber 
		}
	}
}

