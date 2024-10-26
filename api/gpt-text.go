package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client
var apiKey string
var assistantID string

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

func GenerateAsistantTextFromPrompt(prompt string, assistantId string) (string, error) {
	threadID, err := CreateThread()

	// Creates a new conversation in the thread.
	_, err = client.CreateMessage(context.Background(), threadID,
		openai.MessageRequest{
			Role:    "user",
			Content: prompt,
		})

	if err != nil {
		return "", err
	}
	assistantID := os.Getenv(assistantId)

	run, err := client.CreateRun(context.Background(), threadID, openai.RunRequest{
		Model:       "gpt-4o-mini",
		AssistantID: assistantID,
	})

	if err != nil {
		return "", err
	}
	done := false
	tryes := 0
	for !done {
		resp, err := client.RetrieveRun(context.Background(), run.ThreadID, run.ID)
		if err != nil {
			return "", err
		}
		switch resp.Status {
		case openai.RunStatusInProgress:
			continue
		case openai.RunStatusCompleted:
			done = true
		case openai.RunStatusFailed:
			return "", fmt.Errorf("run failed:")
		case openai.RunStatusCancelled:
			return "", fmt.Errorf("run cancelled")
		case openai.RunStatusRequiresAction:
			return "", fmt.Errorf("run requires action")
		default:
			time.Sleep(2 * time.Second)
			tryes++
			if tryes >= 3 {
				return "", fmt.Errorf("Error generando mensaje")
			}
		}
	}

	// Retrieve the most recent message from the thread
	messages, err := client.ListMessage(context.Background(), threadID, nil, nil, nil, nil, nil)

	if err != nil {
		return "", err
	}

	// print the last mssage

	botMessage := messages.Messages[0]

	botSays := botMessage.Content[0].Text.Value

	return botSays, nil

}

func CreateThread() (string, error) {
	ctx := context.Background()

	thread, err := client.CreateThread(ctx, openai.ThreadRequest{
		// it's possible to give a chat history here to continue a conversation
	})
	if err != nil {
		return "", err
	}
	return thread.ID, nil
}
