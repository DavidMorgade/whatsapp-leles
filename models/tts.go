package models

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func CreateTTS(tts string) (string, error) {
	// if audio folder does not exist, create it
	if _, err := os.Stat("public/audio"); os.IsNotExist(err) {
		os.Mkdir("public/audio", 0755)
	}
	speech := htgotts.Speech{Folder: "public/audio", Language: voices.Spanish}

	// create a random number to add to the filename
	// to avoid overwriting files
	randomNumber := rand.Intn(1000)
	randomNumberStr := strconv.Itoa(randomNumber)

	// remove spaces from tts and add it to initial filename
	shortenedTTS := randomNumberStr + strings.ReplaceAll(tts, " ", "")

	if len(shortenedTTS) > 20 {
		shortenedTTS = shortenedTTS[:20]
	}

	// Generate the speech
	fileName, err := speech.CreateSpeechFile(tts, shortenedTTS)

	if err != nil {
		return "", err
	}

	return fileName, nil

}

func SendTTS(audioURL string, client *whatsmeow.Client, v *events.Message) error {
	fmt.Println("Sending audio")

	// Open the image file
	file, err := os.Open(audioURL)
	if err != nil {
		return fmt.Errorf("failed to open audio file: %v", err)
	}
	defer file.Close()

	// Read the image file
	audioData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read audio file: %v", err)
	}

	resp, err := client.Upload(context.Background(), audioData, whatsmeow.MediaAudio)

	if err != nil {
		return fmt.Errorf("failed to upload audio: %v", err)
	}

	// Create the image message
	audioMsg := &waE2E.AudioMessage{
		URL:           &resp.URL, // URL will be filled by WhatsApp server
		Mimetype:      proto.String("audio/mpeg"),
		DirectPath:    &resp.DirectPath,
		FileLength:    &resp.FileLength,
		MediaKey:      resp.MediaKey,
		FileEncSHA256: resp.FileEncSHA256,
		FileSHA256:    resp.FileSHA256,
	}

	fmt.Println(*audioMsg.URL)

	// Send the image message
	_, err = client.SendMessage(context.Background(), v.Info.Chat, &waE2E.Message{
		AudioMessage: audioMsg,
	})
	if err != nil {
		return fmt.Errorf("failed to send audio message: %v", err)
	}

	return nil
}
