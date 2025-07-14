package message

import (
	"encoding/json"
	"time"
)

type MessageType string

const (
	MessageTypePrimeCheck MessageType = "prime_check"
	MessageTypeEmailSend  MessageType = "email_send"
)

type Message struct {
	ID        string          `json:"id"`
	Type      MessageType     `json:"type"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}

type PrimeCheckPayload struct {
	RequestID  int32  `json:"request_id"`
	UserID     int32  `json:"user_id"`
	NumberText string `json:"number_text"`
}

type EmailSendPayload struct {
	RequestID  int32  `json:"request_id"`
	UserID     int32  `json:"user_id"`
	Email      string `json:"email"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	IsPrime    bool   `json:"is_prime"`
	NumberText string `json:"number_text"`
}

func NewMessage(msgType MessageType, payload interface{}) (*Message, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type:      msgType,
		Payload:   payloadBytes,
		CreatedAt: time.Now(),
	}, nil
}

func (m *Message) UnmarshalPrimeCheckPayload() (*PrimeCheckPayload, error) {
	var payload PrimeCheckPayload
	if err := json.Unmarshal(m.Payload, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

func (m *Message) UnmarshalEmailSendPayload() (*EmailSendPayload, error) {
	var payload EmailSendPayload
	if err := json.Unmarshal(m.Payload, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
