package main

import(
	"fmt"
	"strconv"
)

func printRelease(state *gameState){
	fmt.Println("You can only have six pokemon with you.  Pick one to release")
	for i := 1; i < 6; i++{
		fmt.Printf("[%d] %s\n", i, state.pokemon[i].name)
	}
	fmt.Printf("[6] %s\n", state.currentPokemon.name)
}

func processInputRelease(state *gameState, input string){
	i, err := strconv.Atoi(input)
	if err != nil{
		return
	}
	if i < 1 || i > 6{
		return
	}

	if i < 6{
		fmt.Printf("Goodbye %s!\n", state.pokemon[i].name)
		state.pokemon[i] = state.currentPokemon
	}else if i == 6{
		fmt.Printf("Goodbye %s!", state.currentPokemon.name)
	}
	state.scene = Location
}