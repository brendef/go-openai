package main

import (
	"fmt"
	"os"

	"github.com/brendef/go-openai/openai"
)

func main() {

	message := "Please write a short menu description for a cheese burger" // default message

	if os.Args[1] != "" {
		message = os.Args[1]
	}

	ai := openai.NewOpenAi(openai.Config{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.4,
		Context:     "You generate concise menu descriptions using adjectives. Limit responses to 60 characters or less. Keep it 5th-grade level and add South African flavor.",
		MaxTokens:   60,
	})

	response, usedTokens, err := ai.Chat(message)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("respons: ", response)
	fmt.Println("consumed: ", usedTokens, " tokens")

}
