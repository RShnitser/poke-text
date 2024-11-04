package main

import(
	"fmt"
	"strconv"
)

func printTravel(state *gameState){
	fmt.Println("Pick location to explore:")
	for i, loc := range state.locations{
		fmt.Printf("[%d] %s\n", i + 1, loc.name)
	}
	fmt.Printf("[%d] Rest(-1 days)\n", len(state.locations) + 1)
}


func processInputTravel(state *gameState, input string){
	i, err := strconv.Atoi(input)
	if err != nil{
		return
	}
	if i < 1 || i > len(state.locations) + 1{
		return
	}
	if i <= len(state.locations){
		state.currentLocation = i - 1
		state.scene = Location
		fmt.Printf("You travel to the %s\n", state.locations[state.currentLocation].name)
	}else if i == len(state.locations) + 1{
		state.stamina = 20
		state.daysLeft -= 1
		fmt.Println("You rest and recover your stamina")
		if state.daysLeft == 0{
			//state.scene = Time
		}
	}
}