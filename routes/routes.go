package routes

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func RegisterRoutes(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {
	// //ruta solo para pruebas debe estar arriba de las demas
	// if CheckTest(client, v, messageContent) {
	// 	return true
	// }
	// ASSISTANT ROUTES
	if CheckAssistantMention(client, v, messageContent, messageWithoutCommand) {
		return true
	}
	// REGULAR IA ROUTES
	if CheckRegularIAMention(client, v, messageContent, messageWithoutCommand) {
		return true
	}
	// AUDIO ROUTES
	if CheckAudio(client, v, messageContent, messageWithoutCommand) {
		return true
	}
	// IMAGE ROUTES
	if CheckImage(client, v, messageContent, messageWithoutCommand) {
		return true
	}
	// HELP ROUTES
	if CheckHelp(client, v, messageContent) {
		return true
	}
	// CRYPTO ROUTES
	if CheckCrypto(client, v, messageContent, messageWithoutCommand) {
		return true
	}
	// WEATHER ROUTES
	if CheckWeather(client, v, messageContent) {
		return true
	}

	return false

}
