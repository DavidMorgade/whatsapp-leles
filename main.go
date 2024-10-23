package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/whatsapp-leles/db"
	"github.com/whatsapp-leles/utils"
	"github.com/whatsapp-leles/whatsapp"
	"go.mau.fi/whatsmeow"
)

func main() {
	db.CreateDB()
	deviceStore, clientLog, err := db.CreateWaDB()
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(GetEventHandler(client))

	err = utils.CheckWaLogin(client)

	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}

func GetEventHandler(client *whatsmeow.Client) func(interface{}) {

	return func(evt interface{}) {
		whatsapp.CheckMention(client, evt)
	}
}
