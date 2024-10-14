package utils

import (
	"context"
	"log"
	"regexp"
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func CheckBotMention(message *waE2E.ExtendedTextMessage, botId string) bool {
	if message == nil {
		log.Println("CheckBotMention: message is nil")
		return false
	}
	if message.ContextInfo == nil {
		log.Println("CheckBotMention: message.ContextInfo is nil")
		return false
	}
	if message.ContextInfo.MentionedJID == nil {
		log.Println("CheckBotMention: message.ContextInfo.MentionedJID is nil")
		return false
	}
	for _, mention := range message.ContextInfo.MentionedJID {
		if strings.Contains(mention, botId) {
			return true
		}
	}
	return false
}

func CheckIfQuotedMessage(message *waProto.Message) bool {
	if message == nil {
		log.Println("CheckIfQuotedMessage: message is nil")
		return false
	}
	if message.GetExtendedTextMessage() == nil {
		log.Println("CheckIfQuotedMessage: message.GetExtendedTextMessage() is nil")
		return false
	}
	if message.GetExtendedTextMessage().ContextInfo == nil {
		log.Println("CheckIfQuotedMessage: message.GetExtendedTextMessage().ContextInfo is nil")
		return false
	}
	if message.GetExtendedTextMessage().ContextInfo.QuotedMessage == nil {
		log.Println("CheckIfQuotedMessage: message.GetExtendedTextMessage().ContextInfo.QuotedMessage is nil")
		return false
	}
	return true
}

func SendMessage(message string, client *whatsmeow.Client, v *events.Message) {
	if client == nil {
		log.Println("SendMessage: client is nil")
		return
	}
	if v == nil {
		log.Println("SendMessage: event message is nil")
		return
	}
	client.SendMessage(context.Background(), v.Info.Chat, &waProto.Message{
		Conversation: proto.String(message),
	})
}

func RemoveBotId(message string) string {
	re := regexp.MustCompile(`@\S+`)
	return re.ReplaceAllString(message, "")
}
