package message

import (
	"context"
	"encoding/json"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type MessageType string

const (
	MessageTypePrimeCheck MessageType = "prime_check"
	MessageTypeEmailSend  MessageType = "email_send"
)

type Message struct {
	ID           string            `json:"id"`
	Type         MessageType       `json:"type"`
	Payload      json.RawMessage   `json:"payload"`
	CreatedAt    time.Time         `json:"created_at"`
	TraceContext map[string]string `json:"trace_context,omitempty"`
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
	MessageID  string `json:"message_id"`
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

func NewMessageWithTraceContext(ctx context.Context, msgType MessageType, payload interface{}) (*Message, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Extract trace context
	traceContext := make(map[string]string)
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		propagator := otel.GetTextMapPropagator()
		propagator.Inject(ctx, propagation.MapCarrier(traceContext))
	}

	return &Message{
		Type:         msgType,
		Payload:      payloadBytes,
		CreatedAt:    time.Now(),
		TraceContext: traceContext,
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

func (m *Message) ExtractTraceContext(ctx context.Context) context.Context {
	if len(m.TraceContext) == 0 {
		return ctx
	}

	propagator := otel.GetTextMapPropagator()
	return propagator.Extract(ctx, propagation.MapCarrier(m.TraceContext))
}
