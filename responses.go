package turbosms

const (
	// AuthSuccessResponse represents success response string
	AuthSuccessResponse = "Вы успешно авторизировались"
	// MessageSendSuccessResponse represents success response string for a sent message
	MessageSendSuccessResponse = "Сообщения успешно отправлены"
)

// AuthResponse represents auth response from the server
type AuthResponse struct {
	AuthResult string `xml:"AuthResult"`
}

// SendMessageResponse represents sms send response
type SendMessageResponse struct {
	SendSMSResult *ArrayOfString `xml:"SendSMSResult,omitempty"`
}

type ArrayOfString struct {
	//XMLName xml.Name `xml:"http://turbosms.in.ua/api/Turbo ArrayOfString"`
	ResultArray []string `xml:"ResultArray,omitempty"`
}
