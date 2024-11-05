package main

import(
	"fmt"
)

func printStamina(state *gameState){
	fmt.Println("You and your pokemon are too exhausted to continue any further and must leave the area to rest")
	fmt.Println("[1] Continue")
}

func processInputStamina(state *gameState, input string){
	state.locations[state.currentLocation].progress = 0
	state.stamina = 20
	state.daysLeft -= 2

	if state.daysLeft > 0{
		state.scene = Travel
	}else{
		state.scene = Lose
	}
}