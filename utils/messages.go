package utils

import (
	"context"
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func CheckBotMention(message *waE2E.ExtendedTextMessage, botId string) bool {
	for _, mention := range message.ContextInfo.MentionedJID {
		if strings.Contains(mention, botId) {
			return true
		}
	}
	return false
}

func SendMessage(message string, client *whatsmeow.Client, v *events.Message) {
	client.SendMessage(context.Background(), v.Info.Chat, &waProto.Message{
		Conversation: proto.String(message),
	})
}
