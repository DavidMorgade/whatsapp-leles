package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"github.com/whatsapp-leles/db"
	"github.com/whatsapp-leles/models"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	db.CreateDB()
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	dsn := "file:whatsapp.db?_foreign_keys=on"
	container, err := sqlstore.New("sqlite3", dsn, dbLog)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(GetEventHandler(client))

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}

func GetEventHandler(client *whatsmeow.Client) func(interface{}) {

	return func(evt interface{}) {

		switch v := evt.(type) {
		case *events.Message:

			message := v.Message.GetExtendedTextMessage()
			recievedBy := v.Info.PushName
			if message == nil {
				return
			}

			messageModel := models.Message{
				UserID:  1,
				Message: message.GetText(),
			}

			// Check if the message mentions the bot
			botID := client.Store.ID.User
			if utils.CheckBotMention(message, botID) {
				if utils.RemoveBotId(message.GetText()) == " muestra" {
					messages, err := messageModel.GetAllMessages()
					if err != nil {
						fmt.Println(err)
					}
					for _, message := range messages {
						utils.SendMessage("Mensaje guardado de: "+string(message.UserID), client, v)
						utils.SendMessage("Contenido del mensaje: "+message.Message, client, v)
					}
				}
				messageModel.SaveMessage()
				fmt.Printf("Received mention in group: %s\n", utils.RemoveBotId(message.GetText()))
				fmt.Printf("Recieved by: %s\n", recievedBy)

				utils.SendMessage("Soy un bot y estoy funcionando "+recievedBy, client, v)

				break
			}
		}
	}
}
