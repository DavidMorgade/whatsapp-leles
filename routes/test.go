package routes

import (
	"strings"

	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func CheckTest(client *whatsmeow.Client, v *events.Message, messageContent string) bool {

	parsedRoute := strings.ToLower(messageContent)

	//TODO: Add your code here
	if strings.HasPrefix(parsedRoute, " /prueba") {
		utils.SendMessage("PRUEBA", client, v)
		return true
	} else {
		utils.SendMessage("Mi gran amo y señor me tiene en modo mantenimiento, por favor espere a que termine de hacer pruebas", client, v)
		return true
	}

}
