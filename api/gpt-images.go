package api

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

const maxWords = 10

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GenerateImageFromText(prompt string) (string, error) {
	// Ensure the prompt is no longer than 10 words
	words := strings.Fields(prompt)
	if len(words) > maxWords {
		return "", errors.New("prompt exceeds 10 words")
	}

	// Get the OpenAI API key from the environment
	apiKey := os.Getenv("OPEN_AI_KEY")
	if apiKey == "" {
		return "", errors.New("OPEN_AI_KEY is not set in the environment")
	}

	// Create a new OpenAI client
	client := openai.NewClient(apiKey)

	// Create the request payload
	request := openai.ImageRequest{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024", // You can adjust the size as needed
	}

	// Send the request
	response, err := client.CreateImage(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("failed to generate image: %v", err)
	}

	// Return the URL of the generated image
	if len(response.Data) > 0 {
		return response.Data[0].URL, nil
	}

	return "", errors.New("no image URL found in response")
}
