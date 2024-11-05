package main

import(
	"fmt"
)

func printLose(state *gameState){
	fmt.Println("You have taken too long and Mew has left the area")
	fmt.Println("[1] Exit Game")
}

func processInputLose(state *gameState, input string){
	state.running = false
}