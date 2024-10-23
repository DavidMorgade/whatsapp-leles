package whatsapp

import (
	"fmt"
	"github.com/whatsapp-leles/routes"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"regexp"
	"strings"
)

func CheckMention(client *whatsmeow.Client, v any) {
	switch v := v.(type) {
	case *events.Message:

		if !v.Info.IsGroup {
			utils.SendMessage("No se puede utilizar este bot por mensaje privado", client, v)
			return
		}

		message := v.Message.GetExtendedTextMessage()

		messageContent := utils.RemoveBotId(message.GetText())

		// Use a regular expression to remove the command including the /
		re := regexp.MustCompile(`(?i)/\S*`)
		messageWithoutCommand := re.ReplaceAllString(messageContent, "")

		botID := client.Store.ID.User

		if !utils.CheckBotMention(message, botID) {
			break
		}

		if utils.CheckIfQuotedMessage(v.Message) {
			quotedMessage := v.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage().GetConversation()
			fmt.Println("Quoted message: ", quotedMessage)
			if quotedMessage == "" {
				utils.SendMessage("No se puede citar un mensaje vacío o con una imagen", client, v)
				break
			}
			messageWithoutCommand = quotedMessage
		}

		if strings.ReplaceAll(messageContent, " ", "") == "" {
			utils.DefaultHelpMessage(client, v)
			break
		}

		// If the message matches a command, we check with the routes
		if routes.RegisterRoutes(client, v, messageContent, messageWithoutCommand) {
			break
		}

		// Si el mensaje no coincide con ningún comando, se muestra el mensaje de ayuda
		utils.DefaultHelpMessage(client, v)
		break

	}
}
