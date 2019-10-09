package turbosms

import "github.com/tiaguinho/gosoap"

// NewMessage will create new message
// Destination supports many numbers sending example: +380000000999,380000000998
func NewMessage(text, destination, sender string) (message *Message) {
	message = &Message{}
	message.Destination = destination
	message.Sender = sender
	message.Text = text
	return
}

// Message represents the base messages
type Message struct {
	Text        string
	Destination string
	Sender      string
}

// AsParam will return message in required for sending format
func (message *Message) AsParam() (params *gosoap.Params) {
	params = &gosoap.Params{
		"text":        message.Text,
		"destination": message.Destination,
		"sender":      message.Sender,
	}

	return
}
