package repository

import (
	"context"

	"github.com/ponyo877/product-expiry-tracker/db/generated_sql"
	"github.com/ponyo877/product-expiry-tracker/internal/outbox/model"
	"github.com/ponyo877/product-expiry-tracker/internal/outbox/usecase"
)

type OutboxRepository struct {
	queries *generated_sql.Queries
}

func NewOutboxRepository(queries *generated_sql.Queries) usecase.OutboxRepository {
	return &OutboxRepository{
		queries: queries,
	}
}

func (r *OutboxRepository) GetUnprocessedMessages(ctx context.Context) ([]*model.OutboxMessage, error) {
	sqlcMessages, err := r.queries.GetUnprocessedOutboxMessages(ctx)
	if err != nil {
		return nil, err
	}

	messages := make([]*model.OutboxMessage, len(sqlcMessages))
	for i, sqlcMsg := range sqlcMessages {
		messages[i] = model.NewOutboxMessage(
			sqlcMsg.ID,
			sqlcMsg.EventType,
			sqlcMsg.Payload,
			sqlcMsg.Processed,
			sqlcMsg.CreatedAt,
			sqlcMsg.UpdatedAt,
		)
	}

	return messages, nil
}

func (r *OutboxRepository) MarkMessageAsProcessed(ctx context.Context, messageID int32) error {
	return r.queries.MarkOutboxMessageProcessed(ctx, messageID)
}