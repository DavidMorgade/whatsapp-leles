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
		"/ayuda - Muestra la lista de comandos disponibles\n\n" +
		"/tiempo - Muestra el tiempo actual en San Fernando\n\n" +
		"/tiempo [ciudad] - Muestra el tiempo actual en la ciudad especificada\n\n" +
		"/imagen [texto] - Genera una imagen de IA del promp especificado\n\n" +
		"/meme [texto] - Genera un meme con el texto especificado\n\n" +
		"/audio [texto] - Genera una audio con el text especificado"

	SendMessage(message, client, v)
}
