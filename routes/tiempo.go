package routes

import (
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckWeather(client *whatsmeow.Client, v *events.Message, messageContent string) bool {
	parsedRoute := strings.ToLower(messageContent)

	if strings.HasPrefix(parsedRoute, " /tiempo") {
		weather, err := api.GetWeather()
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}

		utils.SendWeatherMessage(*weather, client, v)

		return true
	}

	return false
}
