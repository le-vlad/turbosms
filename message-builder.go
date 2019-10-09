package turbosms

import (
	"fmt"
	"github.com/errors"
)

const (
	ErrMessageBuilderEmptySign        = errors.Error("err. message sign can't be an empty")
	ErrMessageBuilderEmptyDestination = errors.Error("err. message destination can't be an empty")
	ErrMessageBuilderEmptyBody        = errors.Error("err. message body can't be an empty")
)

// NewMessageBuilder will return a new instance of message
func NewMessageBuilder() (messageBuilder *MessageBuilder) {
	messageBuilder = &MessageBuilder{}
	return
}

type MessageBuilder struct {
	sign         string
	destination  string
	body         string
	destinations []string

	singleReceiver bool
}

func (messageBuilder *MessageBuilder) validate() (err error) {
	if messageBuilder.sign == "" {
		err = ErrMessageBuilderEmptySign
		return
	}

	if messageBuilder.body == "" {
		err = ErrMessageBuilderEmptyBody
		return
	}

	if messageBuilder.destination == "" && len(messageBuilder.destinations) == 0 {
		err = ErrMessageBuilderEmptyDestination
		return
	}

	return
}

// SetSign will set a sender name for message
func (messageBuilder *MessageBuilder) SetSign(sign string) *MessageBuilder {
	messageBuilder.sign = sign
	return messageBuilder
}

func (messageBuilder *MessageBuilder) GetSign() string {
	return messageBuilder.sign
}

// SetBody will set a message body content
func (messageBuilder *MessageBuilder) SetBody(body string) *MessageBuilder {
	messageBuilder.body = body
	return messageBuilder
}

func (messageBuilder *MessageBuilder) GetBody() string {
	return messageBuilder.body
}

// SetDestination will set a single receiver for a message
func (messageBuilder *MessageBuilder) SetDestination(phoneNumber string) *MessageBuilder {
	messageBuilder.destination = phoneNumber
	messageBuilder.singleReceiver = true
	return messageBuilder
}

func (messageBuilder *MessageBuilder) GetDestination() string {
	return messageBuilder.destination
}

// SetDestinations will set an array of receivers for a message
func (messageBuilder *MessageBuilder) SetDestinations(phoneNumbers []string) *MessageBuilder {
	messageBuilder.destinations = phoneNumbers
	messageBuilder.singleReceiver = false
	return messageBuilder
}

func (messageBuilder *MessageBuilder) GetDestinations() []string {
	return messageBuilder.destinations
}

// Build will return message instance
func (messageBuilder *MessageBuilder) Build() (message *Message, err error) {
	var (
		destination string
	)

	if err = messageBuilder.validate(); err != nil {
		return
	}

	if messageBuilder.singleReceiver {
		destination = messageBuilder.destination
	} else {
		for _, v := range messageBuilder.destinations {
			destination += fmt.Sprintf("%s,", v)
		}
	}

	message = NewMessage(messageBuilder.body, destination, messageBuilder.sign)
	return
}
