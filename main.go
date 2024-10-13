package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/whatsapp-leles/db"
	"github.com/whatsapp-leles/utils"
	"go.mau.fi/whatsmeow"
	"os"
	"os/signal"
	"syscall"
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
		utils.CheckMention(client, evt)
	}
}
