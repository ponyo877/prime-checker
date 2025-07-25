package usecase

import (
	"context"

	"github.com/ponyo877/prime-checker/internal/outbox/model"
	"github.com/ponyo877/prime-checker/internal/shared/message"
)

type OutboxRepository interface {
	GetUnprocessedMessages(ctx context.Context) ([]*model.OutboxMessage, error)
	MarkMessageAsProcessed(ctx context.Context, messageID int32) error
}

type MessagePublisher interface {
	PublishMessage(ctx context.Context, subject string, msg *message.Message) error
}
