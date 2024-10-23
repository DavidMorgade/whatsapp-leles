package routes

import (
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckRegularIAMention(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {

	parsedRoute := strings.ToLower(messageContent)
	// genera un text con un prompt a la ia
	if strings.HasPrefix(parsedRoute, " /ia") {
		utils.SendMessage("Generando texto...", client, v)
		text, err := api.GenerateAsistantTextFromPrompt(messageWithoutCommand, "ASSISTANT_LELE")
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(text, client, v)
		return true
	}
	// humilla a la persona o echo que se menciona
	if strings.HasPrefix(parsedRoute, " /humillar") {
		utils.SendMessage("Generando texto...", client, v)
		text, err := api.GenerateAsistantTextFromPrompt("En esta respuesta debes humillar a la persona o echo que se menciona, utiliza todas los insultos que encuentres en el archivo de texto para hacerlo "+messageWithoutCommand, "ASSISTANT_LELE")
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(text, client, v)
		return true
	}
	if strings.HasPrefix(parsedRoute, " /alabar") {
		utils.SendMessage("Generando texto...", client, v)
		text, err := api.GenerateAsistantTextFromPrompt("En esta respuesta debes alabar a la persona o echo que se menciona, utiliza todas las palabras posibles que encuentres en el archivo de texto para hacerlo "+messageWithoutCommand, "ASSISTANT_LELE")
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(text, client, v)
		return true
	}
	if strings.HasPrefix(parsedRoute, " /chiste") {
		utils.SendMessage("Generando texto...", client, v)
		text, err := api.GenerateAsistantTextFromPrompt("Cuenta un chiste con las expresiones que usamos en el grupo, y si en el resto de este mensaje aparece algo mas que añadir al chiste añadelo "+messageWithoutCommand, "ASSISTANT_LELE")
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(text, client, v)
		return true
	}

	return false

}
