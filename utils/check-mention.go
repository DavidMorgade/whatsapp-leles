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

		message := v.Message.GetExtendedTextMessage()

		messageContent := RemoveBotId(message.GetText())

		// Use a regular expression to remove the command including the /
		re := regexp.MustCompile(`(?i)/\S*`)
		messageWithoutCommand := re.ReplaceAllString(messageContent, "")

		botID := client.Store.ID.User

		if !CheckBotMention(message, botID) {
			fmt.Println("No se mencionó al bot")
			break
		}

		if CheckIfQuotedMessage(v.Message) {
			quotedMessage := v.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetConversation()
			messageWithoutCommand = quotedMessage
		}

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

		if strings.HasPrefix(strings.ToLower(messageContent), " /audio") {
			SendMessage("Generando audio de"+messageWithoutCommand, client, v)
			audioPath, err := models.CreateTTS(messageWithoutCommand)
			err = models.SendTTS(audioPath, client, v)
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			break
		}
		if strings.HasPrefix(strings.ToLower(messageContent), " /meme") {
			SendMessage("Generando meme de"+messageWithoutCommand, client, v)
			imgURL, err := api.GenerateImageFromText("Genera un meme del siguiente texto" + messageWithoutCommand)
			if err != nil {
				SendMessage("No se pudo generar el meme debido a la cantidad de peticiones, esperate un poquito maquinon", client, v)
			}
			err = SendImage("Meme de:"+messageWithoutCommand, imgURL, client, v)
			if err != nil {
				SendMessage(err.Error(), client, v)
			}
			break
		}
		// Usa inteligencia artificial para generar una imagen a partir de un texto
		if strings.HasPrefix(strings.ToLower(messageContent), " /imagen") {
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
		// Si el mensaje no coincide con ningún comando, se muestra el mensaje de ayuda
		DefaultHelpMessage(client, v)
		break

	}
}
