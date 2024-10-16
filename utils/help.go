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
		"/imagen [texto] - Genera una imagen de IA del promp especificado\n" +
		"/meme [texto] - Genera un meme con el texto especificado\n" +
		"/audio [texto] - Genera una audio con el texto especificado\n" +
		"/ia [texto] - Genera un texto de IA con el prompt especificado\n" +
		"/humillar [texto] - Pongo a parir a la persona que me digas\n" +
		"/chiste [texto] - Cuenta un chiste con el texto especificado\n" +
		"/precio [criptomoneda] - Muestra el precio de la criptomoneda especificada\n\n" +
		"Todos los comandos anteriores funcionan tambien mencionando un comentario + a mi nombre"

	SendMessage(message, client, v)
}
