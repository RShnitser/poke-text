package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

type gameState struct{
	running bool	
}

func (state *gameState)init(){
	state.running = true
}

func (state *gameState)run(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for state.running{
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		fmt.Println(input)
	}
}