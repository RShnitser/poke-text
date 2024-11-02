package main

import (
	"os"
	"fmt"
	//"context"

	 "github.com/joho/godotenv"
	 //openai "github.com/sashabaranov/go-openai"
)

func main(){
	//brightnessArray := "Ñ@#W$9876543210?!abc;:+=-,._ "
	//fmt.Println(brightnessArray)

	err := godotenv.Load(".env")
	if err != nil{
	 fmt.Printf("Error loading .env file: %s", err)
	 os.Exit(1)
	}
   
	//api := os.Getenv("OPENAI_API_KEY")
   
	// client := openai.NewClient(api)
	// resp, err := client.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: "Hello!",
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Printf("ChatCompletion error: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(resp.Choices[0].Message.Content)

	state := gameState{}
	state.init()
	state.run()
}