package utils

import (
	"context"
	"fmt"
	"io"
	"os"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"

	"go.mau.fi/whatsmeow"
	"google.golang.org/protobuf/proto"
)

// SendImage sends an image located at the given URL using the whatsmeow package.
func SendImage(messageContent string, imageURL string, client *whatsmeow.Client, v *events.Message) error {

	// Open the image file
	file, err := os.Open(imageURL)

	if err != nil {
		return fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Read the image file
	imageData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read image file: %v", err)
	}

	resp, err := client.Upload(context.Background(), imageData, whatsmeow.MediaImage)

	if err != nil {
		return fmt.Errorf("failed to upload image: %v", err)
	}

	// Create the image message
	imageMsg := &waE2E.ImageMessage{
		Caption:       proto.String("Generada imagen: " + messageContent),
		URL:           &resp.URL, // URL will be filled by WhatsApp server
		Mimetype:      proto.String("image/jpeg"),
		DirectPath:    &resp.DirectPath,
		FileLength:    &resp.FileLength,
		MediaKey:      resp.MediaKey,
		FileEncSHA256: resp.FileEncSHA256,
		FileSHA256:    resp.FileSHA256,
	}

	fmt.Println(*imageMsg.URL)

	// Send the image message
	_, err = client.SendMessage(context.Background(), v.Info.Chat, &waE2E.Message{
		ImageMessage: imageMsg,
	})
	if err != nil {
		return fmt.Errorf("failed to send image message: %v", err)
	}
	// Delete the image file after sending
	err = os.Remove(imageURL)
	if err != nil {
		return fmt.Errorf("failed to delete image file: %v", err)
	}

	return nil
}
