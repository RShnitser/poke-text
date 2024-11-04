package main

import(
	"fmt"
	"strconv"
	"math/rand"
)

func printCatch(state *gameState){
	pokemon := state.currentPokemon
	if state.escapePercent == 50{
		fmt.Printf("%s is watching you closly\n", pokemon.name)
	}else if state.escapePercent > 50{
		fmt.Printf("%s is angry\n", pokemon.name)
	}else {
		fmt.Printf("%s is eating\n", pokemon.name)
	}

	fmt.Println("[1] Throw rock [-1 stamina]")
	fmt.Println("[2] Throw bait [-1 stamina]")
	fmt.Println("[3] Throw pokeball [-1 stamina]")
	fmt.Println("[4] Escape")
}

func processInputCatch(state *gameState, input string){

	i, err := strconv.Atoi(input)
	if err != nil{
		return
	}
	if i < 1 || i > 4{
		return
	}

	pokemon := state.currentPokemon
	//fmt.Println(pokemon)

	if i == 1{
		state.escapePercent += 5
		state.catchPercent += 5
		state.stamina -= 1
	}else if i == 2{
		state.escapePercent -= 5
		state.catchPercent -= 5
		state.stamina -= 1
	}else if i == 3{
		fmt.Printf("You throw a pokeball at %s\n", pokemon.name)
		state.stamina -= 1
		catch := rand.Intn(100)
		if catch > state.catchPercent{
			fmt.Printf("%s was caught!\n", pokemon.name)
			state.pokemon = append(state.pokemon, pokemon)
			state.scene = Location
			return
		}else{
			fmt.Printf("%s escaped the pokeball!\n", pokemon.name)
		}
	}else if i == 4{
		state.scene = Location
		return
	}
	escape := rand.Intn(100)
	if escape > state.escapePercent{
		fmt.Printf("%s has fled!\n", pokemon.name)
		state.scene = Location
		return
	}

	if state.stamina <= 0 {
		state.scene = Stamina
	}

}