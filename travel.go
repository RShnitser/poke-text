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
}


func processInputTravel(state *gameState, input string){
	i, err := strconv.Atoi(input)
	if err != nil{
		return
	}
	if i < 1 || i > len(state.locations){
		return
	}
	state.currentLocation = i - 1
	state.scene = Location
	fmt.Printf("You travel to the %s\n", state.locations[state.currentLocation].name)
}