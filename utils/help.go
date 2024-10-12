package utils

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func DefaultHelpMessage(client *whatsmeow.Client, v *events.Message) {
	message := "No me has enviado ningún comando válido. Para ver la lista de comandos disponibles, mencioname y escribe /help."
	SendMessage(message, client, v)

}

func SendHelpCommands(client *whatsmeow.Client, v *events.Message) {
	message := "Lista de comandos disponibles:\n\n" +
		"/ayuda - Muestra la lista de comandos disponibles\n" +
		"/tiempo - Muestra el tiempo actual en San Fernando\n" +
		"/tiempo [ciudad] - Muestra el tiempo actual en la ciudad especificada\n" +
		"/muestra - Muestra los mensajes guardados\n" +
		"/guardar - Guarda un mensaje\n"

	SendMessage(message, client, v)
}
