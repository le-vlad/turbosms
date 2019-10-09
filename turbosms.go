package turbosms

import (
	"errors"
	"fmt"

	"github.com/tiaguinho/gosoap"
)

const soapURL = "http://turbosms.in.ua/api/wsdl.html"

// New will return a new instance of TurboSMS client
func New(username, password string) (t *TurboSMS, err error) {
	t = &TurboSMS{}
	t.username = username
	t.password = password

	if err = t.validate(); err != nil {
		return
	}

	if t.soapClient, err = gosoap.SoapClient(soapURL); err != nil {
		err = fmt.Errorf("SoapClient error: %s", err)
	}

	if err = t.auth(); err != nil {
		return
	}

	return
}

// TurboSMS represents the base TurboSMS client
type TurboSMS struct {
	username, password string
	soapClient         *gosoap.Client
}

func (t *TurboSMS) validate() (err error) {
	if t.password == "" {
		return ErrEmptyPassword
	}

	if t.username == "" {
		return ErrEmptyUserName
	}

	return
}

func (t *TurboSMS) auth() (err error) {
	authParam := gosoap.Params{
		"login":    t.username,
		"password": t.password,
	}

	var resp *gosoap.Response
	if resp, err = t.soapClient.Call("Auth", authParam); err != nil {
		return
	}

	var authResponse AuthResponse
	if err = resp.Unmarshal(&authResponse); err != nil {
		return
	}

	if authResponse.AuthResult != AuthSuccessResponse {
		err = errors.New(authResponse.AuthResult)
	}

	return
}

// SendMessage will send a message
func (t *TurboSMS) SendMessage(m *Message) (smsID string, err error) {
	params := m.AsParam()

	var resp *gosoap.Response
	if resp, err = t.soapClient.Call("SendSMS", *params); err != nil {
		return
	}

	var sendMessageResponse SendMessageResponse
	if err = resp.Unmarshal(&sendMessageResponse); err != nil {
		return
	}

	resultCode := sendMessageResponse.SendSMSResult.ResultArray[0]
	if resultCode != MessageSendSuccessResponse {
		err = errors.New(resultCode)
	}

	smsID = sendMessageResponse.SendSMSResult.ResultArray[1]
	return
}
