package main

import (
	"fmt"

	"github.com/brendef/go-openai/openai"
)

func main() {

	ai := openai.NewOpenAi(openai.Config{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
		Context:     "You generate concise menu descriptions using adjectives. Limit responses to 60 characters or less. Keep it 5th-grade level and add South African flavor.",
	})

	response, err := ai.Chat("Please write a short menu description for a cheese burger")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(response)

}
