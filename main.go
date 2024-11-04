package main

func main(){
	var state *gameState = new(gameState)
	state.init()
	state.run()
}