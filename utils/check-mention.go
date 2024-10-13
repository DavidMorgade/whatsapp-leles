package utils

import (
	"fmt"
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/models"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckMention(client *whatsmeow.Client, v any) {
	switch v := v.(type) {
	case *events.Message:

		recievedBy := v.Info.PushName

		message := v.Message.GetExtendedTextMessage()

		messageContent := RemoveBotId(message.GetText())

		botID := client.Store.ID.User
		messageModel := models.Message{
			UserID:  recievedBy,
			Message: messageContent,
		}
		if CheckBotMention(message, botID) {
			if strings.ReplaceAll(messageContent, " ", "") == "" {
				DefaultHelpMessage(client, v)
				break
			}
			if strings.ToLower(messageContent) == " /ayuda" {
				SendHelpCommands(client, v)
				break
			}

			if strings.ToLower(messageContent) == " /tiempo" {
				weather, err := api.GetWeather()
				if err != nil {
					fmt.Println(err)
					SendMessage(err.Error(), client, v)
					break
				}

				SendWeatherMessage(*weather, client, v)

				break
			}

			if strings.HasPrefix(strings.ToLower(messageContent), " /tiempo") {
				city := GetCityFromMessage(messageContent)

				weather, err := api.GetWeatherByCity(city)

				if weather == nil {
					SendMessage("No se encontró la ciudad", client, v)
					break
				}

				fmt.Println("Weather: ", weather)

				if err != nil {
					fmt.Println(err)
					SendMessage(err.Error(), client, v)
					break
				}
				SendWeatherMessage(*weather, client, v)
				break
			}
			if strings.ToLower(messageContent) == " /muestra" {
				messages, err := messageModel.GetAllMessages()
				if err != nil {
					fmt.Println(err)
				}
				for _, message := range messages {
					SendMessage("Mensaje guardado de: "+string(message.UserID), client, v)
					SendMessage("Contenido del mensaje: "+message.Message, client, v)
				}
				break
			}
			if strings.ToLower(messageContent) == " /guarda" {

				messageModel.SaveMessage()
				fmt.Printf("Received mention in group: %s\n", RemoveBotId(message.GetText()))
				fmt.Printf("Recieved by: %s\n", recievedBy)

				break
			}

		}
	}
}
