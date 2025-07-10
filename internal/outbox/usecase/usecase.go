package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ponyo877/product-expiry-tracker/internal/outbox/model"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
)

type OutboxPublishingUsecase struct {
	repo      OutboxRepository
	publisher MessagePublisher
}

func NewOutboxPublishingUsecase(repo OutboxRepository, publisher MessagePublisher) *OutboxPublishingUsecase {
	return &OutboxPublishingUsecase{
		repo:      repo,
		publisher: publisher,
	}
}

func (u *OutboxPublishingUsecase) PublishPendingMessages(ctx context.Context) ([]*model.PublicationResult, error) {
	messages, err := u.repo.GetUnprocessedMessages(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get unprocessed messages: %w", err)
	}

	results := make([]*model.PublicationResult, 0, len(messages))

	for _, outboxMsg := range messages {
		result := u.publishSingleMessage(ctx, outboxMsg)
		results = append(results, result)

		if result.IsSuccess() {
			if err := u.repo.MarkMessageAsProcessed(ctx, result.MessageID()); err != nil {
				log.Printf("Failed to mark message ID %d as processed: %v", result.MessageID(), err)
			}
		}
	}

	return results, nil
}

func (u *OutboxPublishingUsecase) publishSingleMessage(ctx context.Context, outboxMsg *model.OutboxMessage) *model.PublicationResult {
	// Create message for broker
	msg := &message.Message{
		Type:      message.MessageType(outboxMsg.EventType()),
		Payload:   outboxMsg.Payload(),
		CreatedAt: outboxMsg.CreatedAt(),
	}

	subject := u.getSubjectForEventType(outboxMsg.EventType())
	now := time.Now()
	if err := u.publisher.PublishMessage(ctx, subject, msg); err != nil {
		log.Printf("Failed to publish message ID %d: %v", outboxMsg.ID(), err)
		return model.NewPublicationResult(outboxMsg.ID(), model.PublicationStatusFailed, err, now)
	}

	return model.NewPublicationResult(outboxMsg.ID(), model.PublicationStatusSuccess, nil, now)
}

func (u *OutboxPublishingUsecase) getSubjectForEventType(eventType string) string {
	switch eventType {
	case string(message.MessageTypePrimeCheck):
		return "primecheck"
	case string(message.MessageTypeEmailSend):
		return "emailsend"
	default:
		return "unknown"
	}
}
