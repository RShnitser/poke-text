package main

import(
	"fmt"
	"strconv"
	"math/rand"
)

func printLocation(state *gameState){
	loc := state.locations[state.currentLocation]

	fmt.Println(state.data.obstacles[loc.path[loc.progress]].description)
	//name := ""
	cost := 5
	message := "Force your way through"

	Loop:
	for _, pokemon := range state.pokemon{
		for _, ability := range pokemon.abilities{
			if loc.path[loc.progress] == ability{
				cost = 1
				abilityName, ok := state.data.abilityName[ability]
				if !ok{
					abilityName = "UNKNOWN"
				}
				message = fmt.Sprintf("%s uses %s", pokemon.name, abilityName)
				break Loop
			}
		}
	}
	state.currentCost = cost
	fmt.Printf("[1] %s (-%d stamina)\n", message, cost)
	fmt.Printf("[2] Search for pokemon (-1 stamina)\n")
	fmt.Printf("[3] Leave the %s\n", loc.name)
}

func processInputLocation(state *gameState, input string){
	i, err := strconv.Atoi(input)
	if err != nil{
		return
	}
	if i < 1 || i > 3{
		return
	}

	if i == 1{
		state.stamina -= state.currentCost
		if state.stamina <= 0{
			
			state.scene = Stamina
			return
		
		}
		state.locations[state.currentLocation].progress += 1
		if state.locations[state.currentLocation].progress == 5{
			state.locations[state.currentLocation].completed = true
			state.scene = Travel
			fmt.Println("You have found a clue to Mew's location!")
		}
	}else if i == 2{
		data := state.data.locations[state.locations[state.currentLocation].dataIndex]
		rand := rand.Intn(len(data.pokemon))
		pokemon := data.pokemon[rand]
		state.currentPokemon = pokemon
		state.catchPercent = 25
		state.escapePercent = 50
		state.stamina -= 1
		if state.stamina > 0{
			state.scene = Capture
			fmt.Printf("You encounter a wild %s\n", pokemon.name)
		}else{
			state.scene = Stamina
		}
		return
	}else if i == 3{
		state.locations[state.currentLocation].progress += 1
		fmt.Println("exiting")
		state.scene = Travel
		return
	}
}