package routes

import (
	"strings"

	"github.com/whatsapp-leles/models"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckAudio(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {
	parsedRoute := strings.ToLower(messageContent)

	if strings.HasPrefix(parsedRoute, " /audio") {
		utils.SendMessage("Generando audio ...", client, v)
		audioPath, err := models.CreateTTS(messageWithoutCommand)
		err = models.SendTTS(audioPath, client, v)
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		return true
	}

	return false

}
