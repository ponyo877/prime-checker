package repository

import (
	"context"
	"fmt"

	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/model"
	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/usecase"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
)

type ResultPublisher struct {
	outboxRepo usecase.OutboxRepository
}

func NewResultPublisher(outboxRepo usecase.OutboxRepository) usecase.ResultPublisher {
	return &ResultPublisher{
		outboxRepo: outboxRepo,
	}
}

func (p *ResultPublisher) PublishEmailMessage(ctx context.Context, result *model.PrimeResult) error {
	emailPayload := &message.EmailSendPayload{
		RequestID:  result.RequestID(),
		UserID:     result.UserID(),
		Email:      "user@example.com", // TODO: Get from user profile
		Subject:    fmt.Sprintf("Prime Check Result for %s", result.NumberText()),
		Body:       fmt.Sprintf("The number %s is prime: %v", result.NumberText(), result.IsPrime()),
		IsPrime:    result.IsPrime(),
		NumberText: result.NumberText(),
	}

	emailMsg, err := message.NewMessage(message.MessageTypeEmailSend, emailPayload)
	if err != nil {
		return fmt.Errorf("failed to create email message: %w", err)
	}

	return p.outboxRepo.CreateOutboxMessage(ctx, string(message.MessageTypeEmailSend), emailMsg.Payload)
}
