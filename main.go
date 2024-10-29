package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	//brightnessArray := "Ñ@#W$9876543210?!abc;:+=-,._ "
	//fmt.Println(brightnessArray)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		fmt.Println(input)
	}
}