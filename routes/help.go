package routes

import (
	"strings"

	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckHelp(client *whatsmeow.Client, v *events.Message, messageContent string) bool {
	parsedRoute := strings.ToLower(messageContent)

	if strings.HasPrefix(parsedRoute, " /ayuda") {
		utils.SendHelpCommands(client, v)
		return true
	}
	if strings.HasPrefix(parsedRoute, " /version") {
		utils.SendVersionMessage(client, v)
		return true
	}

	return false

}
