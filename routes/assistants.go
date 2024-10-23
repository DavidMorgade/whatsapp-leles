package routes

import (
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckAssistantMention(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {

	parsedRoute := strings.ToLower(messageContent)
	if strings.HasPrefix((parsedRoute), " /toti") {
		utils.SendMessage("Bot toti escribiendo...", client, v)
		text, err := api.GenerateAsistantTextFromPrompt(messageWithoutCommand, "ASSISTANT_TOTI")
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(text, client, v)
		return true
	}

	if strings.HasPrefix(strings.ToLower(messageContent), " /jayn") {
		utils.SendMessage("Bot jayn escribiendo...", client, v)
		text, err := api.GenerateAsistantTextFromPrompt(messageWithoutCommand, "ASSISTANT_JAYN")
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(text, client, v)
		return true
	}

	return false

}
