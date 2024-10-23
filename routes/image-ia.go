package routes

import (
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckImage(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {
	parsedRoute := strings.ToLower(messageContent)

	if strings.HasPrefix(parsedRoute, " /meme") {
		utils.SendMessage("Generando meme...", client, v)
		imgURL, err := api.GenerateImageFromText("Genera un meme del siguiente texto: " + messageWithoutCommand)
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		err = utils.SendImage("Meme de :"+messageWithoutCommand, imgURL, client, v)
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		return true
	}
	// Usa inteligencia artificial para generar una imagen a partir de un texto
	if strings.HasPrefix(parsedRoute, " /ia") {
		utils.SendMessage("Generando imagen...", client, v)
		imgURL, err := api.GenerateImageFromText(messageWithoutCommand)
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		err = utils.SendImage(messageWithoutCommand, imgURL, client, v)
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		return true
	}

	return false
}
