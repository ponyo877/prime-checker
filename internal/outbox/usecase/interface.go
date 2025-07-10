package usecase

import (
	"context"

	"github.com/ponyo877/product-expiry-tracker/internal/outbox/model"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
)

type OutboxRepository interface {
	GetUnprocessedMessages(ctx context.Context) ([]*model.OutboxMessage, error)
	MarkMessageAsProcessed(ctx context.Context, messageID int32) error
}

type MessagePublisher interface {
	PublishMessage(ctx context.Context, subject string, msg *message.Message) error
}
