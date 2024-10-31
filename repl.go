package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Scene int
const (
	Start Scene = iota
	Travel
)

type ability struct{
	name string
}

type pokemon struct{
	name string
	abilities []ability
}

type gameState struct{
	running bool
	stamina int	
	scene Scene
	pokemon pokemon
}

func (state *gameState)init(){
	state.running = true
	state.stamina = 10
	state.scene = Start
}

func getInput(scanner *bufio.Scanner)string{
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)
	return input
}

func printStart(){
	fmt.Println("Pick your partner:")
	fmt.Println("[1] Bulbasaur:")
	fmt.Println("[2] Charmander")
	fmt.Println("[3] Squirtle")
}

func processInputStart(state *gameState, input string){
	
	valid := true
	switch(input){
		case "1":
			state.pokemon = pokemon{"Bulbasaur", nil}
		case "2":
			state.pokemon = pokemon{"Charmander", nil}
		case "3":
			state.pokemon = pokemon{"Squirtle", nil}
		default:
			valid = false
	}
	if valid{
		state.scene = Travel
	}
}

func (state *gameState)run(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for state.running{

		switch(state.scene){
		case Start:
			printStart()
		}
		
		input := getInput(scanner)
		processInputStart(state, input)
		if input == "exit"{
			state.running = false
		}
		//fmt.Println(input)
	}
}