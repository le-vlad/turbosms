package turbosms

import "testing"

var (
	testMessageSign        = "Askmenow"
	testMessageBody        = "Hello, how are you?"
	testMessageDestination = "+380976496914"
)

func TestMessageBuilder_New(t *testing.T) {
	var (
		message *Message
		err     error
	)
	if message, err = buildTestMessage(); err != nil {
		return
	}

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

func TestMessageBuilder_Build(t *testing.T) {
	var (
		m   *Message
		err error
	)

	mb := NewMessageBuilder()
	mb.SetSign(testMessageSign)
	mb.SetDestination(testMessageDestination)
	mb.SetBody(testMessageBody)
	if m, err = mb.Build(); err != nil {
		t.Fatal(err)
	}

	if m == nil {
		t.Fatal("Nil message after build")
	}

}

func TestMessageBuilder_GetBody(t *testing.T) {
	mb := NewMessageBuilder()
	mb.SetBody(testMessageBody)

	if mb.GetBody() != testMessageBody {
		t.Errorf("an error accured. Invalid message body, expected %s ; got %s", testMessageBody, mb.GetBody())
	}

}

func TestMessageBuilder_GetDestination(t *testing.T) {
	mb := NewMessageBuilder()
	mb.SetDestination(testMessageDestination)

	if mb.GetDestination() != testMessageDestination {
		t.Errorf("an error accured. Invalid destinication, expected %s ; got %s", testMessageDestination, mb.GetDestination())
	}

}

func TestMessageBuilder_GetSign(t *testing.T) {
	mb := NewMessageBuilder()
	mb.SetSign(testMessageSign)

	if mb.GetSign() != testMessageSign {
		t.Errorf("an error accured. Invalid message sender, expected %s ; got %s", testMessageDestination, mb.GetSign())
	}
}

func buildTestMessage() (message *Message, err error) {
	m := NewMessageBuilder()
	message, err = m.SetSign(testMessageSign).
		SetBody(testMessageBody).
		SetDestination(testMessageDestination).
		Build()

	return
}
