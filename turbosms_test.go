package turbosms

import (
	"fmt"
	"testing"
)

var (
	username = ""
	password = ""
)

func TestTurboSMS_New(t *testing.T) {
	var (
		err error
	)

	if _, err = New(username, password); err != nil {
		t.Fatal(err)
	}
}

func TestTurboSMS_SendMessage(t *testing.T) {
	var (
		turboSMS  *TurboSMS
		messageID string
		err       error
	)

	if turboSMS, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	message, _ := buildTestMessage()
	if messageID, err = turboSMS.SendMessage(message); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Message was sent. ID: %s", messageID)
}
