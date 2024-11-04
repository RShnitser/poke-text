package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
)

type Scene int
const (
	Start Scene = iota
	Travel
	Location
	Capture
	Stamina
	Time
)

// type Ability string
// const (
// 	Cut Ability = "Cut"
// 	Surf Ability = "Surf"
// 	Strength Ability = "Strength"
// 	Flash Ability = "Flash"
// 	Defog Ability = "Defog"
// 	RockSmash Ability = "Rock Smash"
// 	Climb Ability = "Climb"
// 	Fly Ability = "Fly"
// 	Whirlpool Ability = "Whirlpool"
// 	Waterfall Ability = "Waterfall"
// 	Dive Ability = "Dive"
// )

type Ability int
const (
	Cut Ability = iota
	Surf 
	Strength 
	Flash 
	Defog 
	RockSmash 
	Climb 
	Fly 
	Whirlpool 
	Waterfall 
	Dive
)

type Obstacle struct{
	description string
	solution Ability
}

//type obstacleHandle int
type LocationData struct{
	name string
	//completed bool
	obstacles []Ability
	pokemon []pokemon
}


// type ability struct{
// 	name string
// }

type pokemon struct{
	name string
	abilities []Ability
}

type gameData struct{
	abilityName map[Ability]string
	obstacles []Obstacle
	locations []LocationData
}

//type obstacleHandle int
type LocationState struct{
	dataIndex int
	name string
	completed bool
	progress int
	path []Ability
}

type gameState struct{ 
	data *gameData
	running bool
	stamina int	
	scene Scene
	daysLeft int
	pokemon []pokemon
	locations []LocationState
	currentLocation int
	currentCost int	
	currentPokemon pokemon
	catchPercent int
	escapePercent int
}

func (state *gameState)init(){
	var data *gameData = new(gameData)

	data.abilityName = make(map[Ability]string)
	data.abilityName[Cut] = "Cut"
	data.abilityName[Surf] = "Surf"
	data.abilityName[Strength] = "Strength"
	data.abilityName[Flash] = "Flash"
	data.abilityName[Defog] = "Defog"
	data.abilityName[RockSmash] = "Rock Smash"
	data.abilityName[Climb] = "Climb"
	data.abilityName[Fly] = "Fly"
	data.abilityName[Whirlpool] = "Whirlpool"
	data.abilityName[Waterfall] = "Waterfall"
	data.abilityName[Dive] = "Dive"

	data.obstacles = []Obstacle{}
	data.obstacles = append(data.obstacles, Obstacle{"A small tree blocks your path", Cut})
	data.obstacles = append(data.obstacles, Obstacle{"You come across a still body of water", Surf})
	data.obstacles = append(data.obstacles, Obstacle{"There is a rock blocking the way, but it looks like it can be moved with enough effort", Strength})
	data.obstacles = append(data.obstacles, Obstacle{"The area you come across is dark, and it is very difficult to naviagate", Flash})
	data.obstacles = append(data.obstacles, Obstacle{"The thick fog makes it very difficult to move further", Defog})
	data.obstacles = append(data.obstacles, Obstacle{"Rocks in your way prevent you from moving further, but it looks like they can be broken", RockSmash})
	data.obstacles = append(data.obstacles, Obstacle{"You will have to climb if you want to go further", Climb})
	data.obstacles = append(data.obstacles, Obstacle{"You come across a crevice to wide to jump over", Fly})
	data.obstacles = append(data.obstacles, Obstacle{"A raging whirlpool in the water is in your way", Whirlpool})
	data.obstacles = append(data.obstacles, Obstacle{"You come across a waterfall", Waterfall})
	data.obstacles = append(data.obstacles, Obstacle{"To proceed further will require going underwater", Dive})
	
	data.locations = []LocationData{}
	cave := LocationData{"Cave", []Ability{Surf, Strength, Flash, Dive, Whirlpool, Climb}, []pokemon{}}
	cave.pokemon = append(cave.pokemon, pokemon{"Zubat", []Ability{Defog}})
	cave.pokemon = append(cave.pokemon, pokemon{"Geodude", []Ability{Strength}})
	cave.pokemon = append(cave.pokemon, pokemon{"Onyx", []Ability{Climb}})
	data.locations = append(data.locations, cave)

	forest := LocationData{"Forest", []Ability{Cut, Flash, Defog, Climb, RockSmash}, []pokemon{}}
	forest.pokemon = append(forest.pokemon, pokemon{"Pikachu", []Ability{Flash}})
	forest.pokemon = append(forest.pokemon, pokemon{"Bellsprout", []Ability{Cut}})
	forest.pokemon = append(forest.pokemon, pokemon{"Onyx", []Ability{RockSmash}})
	data.locations = append(data.locations, forest)

	grass := LocationData{"Tall Grass", []Ability{Cut, Defog, Fly, Strength, Climb}, []pokemon{}}
	grass.pokemon = append(grass.pokemon, pokemon{"Pidgey", []Ability{Fly}})
	grass.pokemon = append(grass.pokemon, pokemon{"Rattata", []Ability{Cut}})
	grass.pokemon = append(grass.pokemon, pokemon{"Psyduck", []Ability{Surf}})
	data.locations = append(data.locations, grass)

	coast := LocationData{"Coast", []Ability{Surf, Dive, Whirlpool, Waterfall, RockSmash}, []pokemon{}}
	coast.pokemon = append(coast.pokemon, pokemon{"Goldeen", []Ability{Waterfall}})
	coast.pokemon = append(coast.pokemon, pokemon{"Seel", []Ability{Dive}})
	coast.pokemon = append(coast.pokemon, pokemon{"Shellder", []Ability{Whirlpool}})
	data.locations = append(data.locations, coast)

	state.data = data
	state.running = true
	state.stamina = 20
	state.daysLeft = 10
	state.scene = Start
	state.pokemon = []pokemon{}

	for i, loc := range data.locations{
		locState := LocationState{}
		locState.dataIndex = i
		locState.progress = 0
		locState.completed = false
		locState.name = loc.name
		locState.path = []Ability{}
		//locState.path = []int{}

		for i := 0; i < 5; i++{
			randomObstacle := rand.Intn(len(loc.obstacles))
			locState.path = append(locState.path, loc.obstacles[randomObstacle])
			//location.path =  append(locState.path, loc.obstacles[randomObstacle])
		}
		state.locations = append(state.locations, locState)
	}
}

func getInput(scanner *bufio.Scanner)string{
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)
	return input
}

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



func printState(state *gameState){
	switch(state.scene){
	case Start:
		printStart()
	case Travel:
		printTravel(state)
	case Location:
		printLocation(state)
	case Capture:
		printCatch(state)
	}
}

func processInput(state *gameState, input string){
	switch(state.scene){
	case Start:
		processInputStart(state, input)
	case Travel:
		processInputTravel(state, input)
	case Location:
		processInputLocation(state, input)
	case Capture:
		processInputCatch(state, input)
	}
}

func (state *gameState)run(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for state.running{

		printState(state)
		input := getInput(scanner)
		processInput(state, input)
		if input == "exit"{
			state.running = false
			return
		}
		//fmt.Println(input)
	}
}