package turbosms

import (
	"testing"
)

func TestMessage_New(t *testing.T) {
	message := NewMessage(testMessageBody, testMessageDestination, testMessageSign)
	if message.Destination != testMessageDestination {
		t.Errorf("an error accured. Invalid destinication, expected %s ; got %s", testMessageDestination, message.Destination)
	}

	if message.Text != testMessageBody {
		t.Errorf("an error accured. Invalid message body, expected %s ; got %s", testMessageBody, message.Text)
	}

	if message.Sender != testMessageSign {
		t.Errorf("an error accured. Invalid sender sign, expected %s ; got %s", testMessageSign, message.Sender)
	}
}

func TestMessage_AsParam(t *testing.T) {
	message, _ := buildTestMessage()
	messageParams := message.AsParam()
	messageParamsUnnamed := *messageParams
	if messageParamsUnnamed["destination"] != testMessageDestination {
		t.Errorf("an error accured. Invalid destinication, expected %s ; got %s", testMessageDestination, message.Destination)
	}

	if messageParamsUnnamed["text"] != testMessageBody {
		t.Errorf("an error accured. Invalid message body, expected %s ; got %s", testMessageBody, message.Text)
	}

	if messageParamsUnnamed["sender"] != testMessageSign {
		t.Errorf("an error accured. Invalid sender sign, expected %s ; got %s", testMessageSign, message.Sender)
	}
}
