package routes

import (
	"strings"

	"github.com/whatsapp-leles/api"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckCrypto(client *whatsmeow.Client, v *events.Message, messageContent string, messageWithoutCommand string) bool {
	parsedRoute := strings.ToLower(messageContent)

	if strings.HasPrefix(parsedRoute, " /precio") {
		cryptoInfo, err := api.GetCryptoPrice(messageWithoutCommand)
		if err != nil {
			utils.SendMessage(err.Error(), client, v)
			return true
		}
		utils.SendMessage(cryptoInfo, client, v)
		return true
	}

	return false
}
