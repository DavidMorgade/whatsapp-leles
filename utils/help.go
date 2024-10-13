package utils

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func DefaultHelpMessage(client *whatsmeow.Client, v *events.Message) {
	message := "No me has enviado ningún comando válido. Para ver la lista de comandos disponibles, mencioname y escribe /ayuda."
	SendMessage(message, client, v)

}

func SendHelpCommands(client *whatsmeow.Client, v *events.Message) {
	message := "Lista de comandos disponibles:\n\n" +
		"/ayuda - Muestra la lista de comandos disponibles\n" +
		"/tiempo - Muestra el tiempo actual en San Fernando\n" +
		"/tiempo [ciudad] - Muestra el tiempo actual en la ciudad especificada\n" +
		"/generar [texto] - Genera una imagen de IA del promp especificado\n" +
		"/generarmeme [texto] - Genera un meme con el texto especificado\n"

	SendMessage(message, client, v)
}
