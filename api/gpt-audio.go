package api

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("OPEN_AI_KEY")
	if apiKey == "" {
		fmt.Println("OPEN_AI_KEY is not set in the environment")
	}

	client = openai.NewClient(apiKey)

}

func GenerateAudioFromText(prompt string) {

	apiKey := os.Getenv("OPEN_AI_KEY")

	if apiKey == "" {
		fmt.Println("OPEN_AI_KEY is not set in the environment")
	}

	client := openai.NewClient(apiKey)

	request := openai.CreateSpeechRequest{
		Model: "tts-1",
		Voice: "onyx",
		Input: prompt,
	}

	response, err := client.CreateSpeech(context.Background(), request)

	if err != nil {
		fmt.Println("Error generating audio: ", err)
	}

	fmt.Println(response)

}
