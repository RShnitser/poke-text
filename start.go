package main

import(
	"fmt"
)

func printStart(){
	fmt.Println("Pick your partner:")
	fmt.Println("[1] Bulbasaur")
	fmt.Println("[2] Charmander")
	fmt.Println("[3] Squirtle")
}

func processInputStart(state *gameState, input string){
	
	valid := true
	switch(input){
		case "1":
			state.pokemon = append(state.pokemon, pokemon{"Bulbasaur", []Ability{Cut, RockSmash}})
		case "2":
			state.pokemon = append(state.pokemon, pokemon{"Charmander", []Ability{Strength, Flash}})
		case "3":
			state.pokemon = append(state.pokemon, pokemon{"Squirtle", []Ability{Surf, Dive}})
		default:
			valid = false
	}
	if valid{
		state.scene = Travel
	}
}