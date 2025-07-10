package adapter

import (
	"context"

	"github.com/ponyo877/product-expiry-tracker/internal/outbox/usecase"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/infrastructure"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
)

type MessagePublisher struct {
	broker infrastructure.MessageBroker
}

func NewMessagePublisher(broker infrastructure.MessageBroker) usecase.MessagePublisher {
	return &MessagePublisher{
		broker: broker,
	}
}

func (p *MessagePublisher) PublishMessage(ctx context.Context, subject string, msg *message.Message) error {
	return p.broker.Publish(ctx, subject, msg)
}
