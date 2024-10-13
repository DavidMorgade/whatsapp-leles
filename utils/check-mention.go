package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/models"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckMention(client *whatsmeow.Client, v any) {
	switch v := v.(type) {
	case *events.Message:

		if !v.Info.IsGroup {
			SendMessage("No se puede utilizar este bot por mensaje privado", client, v)
			return
		}

		recievedBy := v.Info.PushName

		message := v.Message.GetExtendedTextMessage()

		messageContent := RemoveBotId(message.GetText())

		// Use a regular expression to remove the command including the /
		re := regexp.MustCompile(`(?i)/\S*`)
		messageWithoutCommand := re.ReplaceAllString(messageContent, "")

		botID := client.Store.ID.User
		messageModel := models.Message{
			UserID:  recievedBy,
			Message: messageContent,
		}
		// Si el mensaje contiene una mención al bot vacía, se muestra el mensaje de ayuda
		if CheckBotMention(message, botID) {
			if strings.ReplaceAll(messageContent, " ", "") == "" {
				DefaultHelpMessage(client, v)
				break
			}
			// Muestra los comandos disponibles
			if strings.ToLower(messageContent) == " /ayuda" {
				SendHelpCommands(client, v)
				break
			}

			// Usa la API de OpenWeatherMap para obtener el tiempo en una ciudad
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
			// Usa inteligencia artificial para generar una imagen a partir de un texto
			if strings.HasPrefix(strings.ToLower(messageContent), " /generar") {
				SendMessage("Generando imagen de"+messageWithoutCommand, client, v)
				imgURL, err := api.GenerateImageFromText(messageWithoutCommand)
				if err != nil {
					SendMessage("No se pudo generar la imagen debido a la cantidad de peticiones, esperate un poquito maquinon", client, v)
					break
				}
				err = SendImage(messageWithoutCommand, imgURL, client, v)
				if err != nil {
					SendMessage(err.Error(), client, v)
					break
				}
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
