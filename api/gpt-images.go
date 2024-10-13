package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
		Size:   "256x256", // You can adjust the size as needed
	}

	// Send the request
	response, err := client.CreateImage(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("failed to generate image: %v", err)
	}

	// Check if the response contains image data
	if len(response.Data) == 0 {
		return "", errors.New("no image URL found in response")
	}

	// Get the image URL
	imageURL := response.Data[0].URL

	// Download the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// Create the file
	fileName := strings.ReplaceAll(prompt, " ", "") + ".png"
	publicPath := filepath.Join("public", "images")

	err = os.MkdirAll(publicPath, os.ModePerm)

	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}
	file, err := os.Create(filepath.Join(publicPath, fileName))
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	filePath := filepath.Join(publicPath, fileName)
	defer file.Close()

	// Save the image to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	// Return the file path
	return filePath, nil
}
