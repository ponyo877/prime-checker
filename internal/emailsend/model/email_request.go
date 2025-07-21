package model

import "time"

type EmailRequest struct {
	requestID  int32
	userID     int32
	email      string
	subject    string
	body       string
	isPrime    bool
	numberText string
	messageID  string
	timestamp  time.Time
}

func NewEmailRequest(requestID, userID int32, email, subject, body string, isPrime bool, numberText, messageID string) *EmailRequest {
	return &EmailRequest{
		requestID:  requestID,
		userID:     userID,
		email:      email,
		subject:    subject,
		body:       body,
		isPrime:    isPrime,
		numberText: numberText,
		messageID:  messageID,
		timestamp:  time.Now(),
	}
}

func (e *EmailRequest) RequestID() int32 {
	return e.requestID
}

func (e *EmailRequest) UserID() int32 {
	return e.userID
}

func (e *EmailRequest) Email() string {
	return e.email
}

func (e *EmailRequest) Subject() string {
	return e.subject
}

func (e *EmailRequest) Body() string {
	return e.body
}

func (e *EmailRequest) IsPrime() bool {
	return e.isPrime
}

func (e *EmailRequest) NumberText() string {
	return e.numberText
}

func (e *EmailRequest) MessageID() string {
	return e.messageID
}

func (e *EmailRequest) Timestamp() time.Time {
	return e.timestamp
}
