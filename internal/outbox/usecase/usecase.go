package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/ponyo877/prime-checker/internal/outbox/model"
	"github.com/ponyo877/prime-checker/internal/shared/message"
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
	tracer := otel.Tracer("outbox-publisher")
	ctx, span := tracer.Start(ctx, "PublishSingleMessage")
	defer span.End()

	// Deserialize the stored message (which includes trace context)
	var msg message.Message
	if err := json.Unmarshal(outboxMsg.Payload(), &msg); err != nil {
		log.Printf("Failed to unmarshal message ID %d: %v", outboxMsg.ID(), err)
		now := time.Now()
		return model.NewPublicationResult(outboxMsg.ID(), model.PublicationStatusFailed, err, now)
	}

	// Extract trace context and continue the trace
	if msg.TraceContext != nil && len(msg.TraceContext) > 0 {
		ctx = msg.ExtractTraceContext(ctx)
		span.End() // End the current span
		ctx, span = tracer.Start(ctx, "PublishSingleMessage") // Start a new span with the extracted context
		defer span.End()
		
		traceID := span.SpanContext().TraceID().String()
		log.Printf("Publishing message ID %d with Trace ID: %s", outboxMsg.ID(), traceID)
	}

	subject := u.getSubjectForEventType(outboxMsg.EventType())
	now := time.Now()
	if err := u.publisher.PublishMessage(ctx, subject, &msg); err != nil {
		span.RecordError(err)
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
