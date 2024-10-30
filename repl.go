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
}

func (state *gameState)init(){
	state.running = true
	state.stamina = 10
	state.scene = Start
}

func (state *gameState)run(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for state.running{

		switch(state.scene){
		case Start:
			fmt.Println("Pick your partner:")
			fmt.Println("[1] Bulbasaur:")
			fmt.Println("[2] Charmander")
			fmt.Println("[3] Squirtle")
		}
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)

		if input == "exit"{
			state.running = false
		}
		//fmt.Println(input)
	}
}