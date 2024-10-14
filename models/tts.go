package models

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

// SplitText splits the input text into chunks of 200 characters or less
func SplitText(text string, chunkSize int) []string {
	var chunks []string
	for len(text) > chunkSize {
		chunks = append(chunks, text[:chunkSize])
		text = text[chunkSize:]
	}
	chunks = append(chunks, text)
	return chunks
}

func CreateTTS(tts string) (string, error) {
	// if audio folder does not exist, create it
	if _, err := os.Stat("public/audio"); os.IsNotExist(err) {
		os.Mkdir("public/audio", 0755)
	}
	speech := htgotts.Speech{Folder: "public/audio", Language: voices.Spanish}

	// Split the input text into chunks
	chunks := SplitText(tts, 200)

	var audioFiles []string
	for _, chunk := range chunks {
		// create a random number to add to the filename to avoid overwriting files
		randomNumber := rand.Intn(1000)
		randomNumberStr := strconv.Itoa(randomNumber)

		// remove spaces from chunk and add it to initial filename
		shortenedChunk := randomNumberStr + strings.ReplaceAll(chunk, " ", "")
		if len(shortenedChunk) > 20 {
			shortenedChunk = shortenedChunk[:20]
		}

		// Generate the speech for each chunk
		fileName, err := speech.CreateSpeechFile(chunk, shortenedChunk)
		if err != nil {
			return "", err
		}
		audioFiles = append(audioFiles, fileName)
	}

	// Concatenate the audio files
	finalFileName := "public/audio/final_" + strconv.Itoa(rand.Intn(100000)) + ".mp3"
	err := ConcatenateAudioFiles(audioFiles, finalFileName)
	if err != nil {
		return "", err
	}

	// Remove the individual audio files
	for _, file := range audioFiles {
		err := os.Remove(file)
		if err != nil {
			return "", fmt.Errorf("failed to remove file %s: %v", file, err)
		}
	}

	return finalFileName, nil
}

// ConcatenateAudioFiles concatenates multiple audio files into a single file
func ConcatenateAudioFiles(files []string, output string) error {
	cmd := exec.Command("ffmpeg", append([]string{"-y", "-i", "concat:" + strings.Join(files, "|"), "-c", "copy", output})...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to concatenate audio files: %v", err)
	}
	return nil
}

func SendTTS(audioURL string, client *whatsmeow.Client, v *events.Message) error {
	fmt.Println("Sending audio")

	// Open the audio file
	file, err := os.Open(audioURL)
	if err != nil {
		return fmt.Errorf("failed to open audio file: %v", err)
	}
	defer file.Close()

	// Read the audio file
	audioData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read audio file: %v", err)
	}

	resp, err := client.Upload(context.Background(), audioData, whatsmeow.MediaAudio)
	if err != nil {
		return fmt.Errorf("failed to upload audio: %v", err)
	}

	// Create the audio message
	audioMsg := &waE2E.AudioMessage{
		URL:           &resp.URL,
		Mimetype:      proto.String("audio/mpeg"),
		DirectPath:    &resp.DirectPath,
		FileLength:    &resp.FileLength,
		MediaKey:      resp.MediaKey,
		FileEncSHA256: resp.FileEncSHA256,
		FileSHA256:    resp.FileSHA256,
	}

	fmt.Println(*audioMsg.URL)

	// Send the audio message
	_, err = client.SendMessage(context.Background(), v.Info.Chat, &waE2E.Message{
		AudioMessage: audioMsg,
	})
	if err != nil {
		return fmt.Errorf("failed to send audio message: %v", err)
	}

	return nil
}
