package main

import(
	"fmt"
	"strconv"
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
	fmt.Printf("[2] Leave the%s\n", loc.name)
}

func processInputLocation(state *gameState, input string){
	i, err := strconv.Atoi(input)
	if err != nil{
		return
	}
	if i < 1 || i > 2{
		return
	}

	if i == 1{
		state.stamina -= state.currentCost
		if state.stamina <= 0{
			state.locations[state.currentLocation].progress = 0
			fmt.Println("You and your pokemon are too exhausted to continue any further and must leave the area to rest")
			state.stamina = 20
			state.daysLeft -= 2
			if state.daysLeft <= 0{
				fmt.Println("You have taken too long and Mew has left the area")
				return
			}
			//fmt.Println("")

		}
		state.locations[state.currentLocation].progress += 1
		if state.locations[state.currentLocation].progress == 5{
			state.locations[state.currentLocation].completed = true
			state.scene = Travel
			fmt.Println("You have found a clue to Mew's location!")
		}
	}
}