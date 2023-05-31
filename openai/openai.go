package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	COMPLETIONS_URL = "https://api.openai.com/v1/chat/completions"
)

type OpenAi struct {
	apiKey       string
	organisation string
	model        string
	temperature  float64
	context      string
}

type Config struct {
	Model       string
	Temperature float64
	Context     string
}

func NewOpenAi(config Config) *OpenAi {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	key := os.Getenv("OPENAI_API_KEY")
	org := os.Getenv("OPENAI_ORGANISATION")

	return &OpenAi{
		apiKey:       key,
		organisation: org,
		model:        config.Model,
		temperature:  config.Temperature,
		context:      config.Context,
	}

}

func (ai *OpenAi) request(method string, url string, body io.Reader) (response *http.Response, err error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {

		return
	}

	req.Header.Add("Authorization", "Bearer "+ai.apiKey)
	req.Header.Add("Content-Type", "application/json")

	if ai.organisation != "" {
		req.Header.Add("OpenAI-Organization", ai.organisation)
	}

	client := &http.Client{}
	response, err = client.Do(req)

	return
}

func (ai *OpenAi) post(url string, input any) (response []byte, err error) {

	respJson, err := json.Marshal(input)
	if err != nil {

		return
	}

	respBytes, err := ai.request(http.MethodPost, url, bytes.NewReader(respJson))
	if err != nil {

		return
	}

	defer respBytes.Body.Close()

	response, err = io.ReadAll(respBytes.Body)

	return
}

func (ai *OpenAi) completion(req ChatRequest) (res ChatResponse, err error) {
	resBytes, err := ai.post(COMPLETIONS_URL, req)
	if err != nil {

		return
	}

	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		fmt.Println(err)

		return
	}

	return
}

func (ai *OpenAi) Chat(message string) (text string, err error) {

	var resp ChatResponse
	resp, err = ai.completion(ChatRequest{
		Model: ai.model,
		Messages: []Message{
			{
				Role:    "system",
				Content: ai.context,
			},
			{
				Role:    "user",
				Content: message,
			},
		},
		Temperature: ai.temperature,
	})

	if err != nil {
		fmt.Println(err)
	}

	text = resp.Choices[0].Text

	return
}
