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
			fmt.Println("Quoted message: ", quotedMessage)
			if quotedMessage == "" {
				SendMessage("No se puede citar un mensaje vacío o con una imagen", client, v)
				break
			}
			messageWithoutCommand = quotedMessage
		}

		if strings.ReplaceAll(messageContent, " ", "") == "" {
			DefaultHelpMessage(client, v)
			break
		}

		// SOLO PARA PRUEBAS MANTENER COMENTADO////////////////////////
		// if strings.HasPrefix(strings.ToLower(messageContent), " /prueba") {
		// 	SendMessage("Generando imagen...", client, v)
		// 	imgURL, err := api.GenerateImageFromText(messageWithoutCommand)
		// 	if err != nil {
		// 		SendMessage(err.Error(), client, v)
		// 		break
		// 	}
		// 	err = SendImage(messageWithoutCommand, imgURL, client, v)
		// 	if err != nil {
		// 		SendMessage(err.Error(), client, v)
		// 		break
		// 	}
		// 	break
		// } else {
		// 	SendMessage("Mi gran amo y señor me tiene en modo mantenimiento, por favor espere a que termine de hacer pruebas", client, v)
		// 	break
		// }
		///////////////////////////SOLO PARA PRUEBAS MANTENER COMENTADO////////////////////////
		if strings.HasPrefix(strings.ToLower(messageContent), " /precio") {
			cryptoInfo, err := api.GetCryptoPrice(messageWithoutCommand)
			if err != nil {
				SendMessage(err.Error(), client, v)
			}
			SendMessage(cryptoInfo, client, v)
			break
		}
		// ASISTENTS personalizados
		if strings.HasPrefix(strings.ToLower(messageContent), " /toti") {
			SendMessage("Bot toti escribiendo...", client, v)
			text, err := api.GenerateAsistantTextFromPrompt(messageWithoutCommand, "ASSISTANT_TOTI")
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			SendMessage(text, client, v)
			break
		}

		if strings.HasPrefix(strings.ToLower(messageContent), " /jayn") {
			SendMessage("Bot jayn escribiendo...", client, v)
			text, err := api.GenerateAsistantTextFromPrompt(messageWithoutCommand, "ASSISTANT_JAYN")
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			SendMessage(text, client, v)
			break
		}
		// genera un text con un prompt a la ia
		if strings.HasPrefix(strings.ToLower(messageContent), " /ia") {
			SendMessage("Generando texto...", client, v)
			text, err := api.GenerateAsistantTextFromPrompt(messageWithoutCommand, "ASSISTANT_LELE")
			if err != nil {
				SendMessage(err.Error(), client, v)
			}
			SendMessage(text, client, v)
			break
		}
		// humilla a la persona o echo que se menciona
		if strings.HasPrefix(strings.ToLower(messageContent), " /humillar") {
			SendMessage("Generando texto...", client, v)
			text, err := api.GenerateAsistantTextFromPrompt("En esta respuesta debes humillar a la persona o echo que se menciona, utiliza todas los insultos que encuentres en el archivo de texto para hacerlo "+messageWithoutCommand, "ASSISTANT_LELE")
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			SendMessage(text, client, v)
			break
		}
		if strings.HasPrefix(strings.ToLower(messageContent), " /alabar") {
			SendMessage("Generando texto...", client, v)
			text, err := api.GenerateAsistantTextFromPrompt("En esta respuesta debes alabar a la persona o echo que se menciona, utiliza todas las palabras posibles que encuentres en el archivo de texto para hacerlo "+messageWithoutCommand, "ASSISTANT_LELE")
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			SendMessage(text, client, v)
			break
		}
		if strings.HasPrefix(strings.ToLower(messageContent), " /chiste") {
			SendMessage("Generando texto...", client, v)
			text, err := api.GenerateAsistantTextFromPrompt("Cuenta un chiste con las expresiones que usamos en el grupo, y si en el resto de este mensaje aparece algo mas que añadir al chiste añadelo "+messageWithoutCommand, "ASSISTANT_LELE")
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			SendMessage(text, client, v)
			break
		}
		// Muestra los comandos disponibles
		if strings.ToLower(messageContent) == " /ayuda" {
			SendHelpCommands(client, v)
			break
		}
		if strings.ToLower(messageContent) == " /version" {
			SendVersionMessage(client, v)
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
			SendMessage("Generando audio ...", client, v)
			audioPath, err := models.CreateTTS(messageWithoutCommand)
			err = models.SendTTS(audioPath, client, v)
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			break
		}
		if strings.HasPrefix(strings.ToLower(messageContent), " /meme") {
			SendMessage("Generando meme...", client, v)
			imgURL, err := api.GenerateImageFromText("Genera un meme del siguiente texto: " + messageWithoutCommand)
			if err != nil {
				SendMessage(err.Error(), client, v)
				break
			}
			err = SendImage("Meme de :"+messageWithoutCommand, imgURL, client, v)
			if err != nil {
				SendMessage(err.Error(), client, v)
			}
			break
		}
		// Usa inteligencia artificial para generar una imagen a partir de un texto
		if strings.HasPrefix(strings.ToLower(messageContent), " /imagen") {
			SendMessage("Generando imagen...", client, v)
			imgURL, err := api.GenerateImageFromText(messageWithoutCommand)
			if err != nil {
				SendMessage(err.Error(), client, v)
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
