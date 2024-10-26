package routes

import (
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckIAAudio(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {
	parsedRoute := strings.ToLower(messageContent)

	if strings.HasPrefix(parsedRoute, " /onyx") {
		utils.SendMessage("Generando audio...", client, v)
		api.GenerateAudioFromText(messageWithoutCommand)

		return true
	}

	return false
}
