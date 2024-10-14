package models

import (
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

func CreateTTS(tts string) (string, error) {
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

	fullPath := filepath.Join("public/audio", fileName)

	return fullPath, nil

}
