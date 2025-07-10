package repository

import (
	"context"
	"encoding/json"

	"github.com/ponyo877/product-expiry-tracker/db/generated_sql"
	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/usecase"
)

type OutboxRepository struct {
	queries *generated_sql.Queries
}

func NewOutboxRepository(queries *generated_sql.Queries) usecase.OutboxRepository {
	return &OutboxRepository{
		queries: queries,
	}
}

func (r *OutboxRepository) CreateOutboxMessage(ctx context.Context, eventType string, payload json.RawMessage) error {
	_, err := r.queries.CreateOutboxMessage(ctx, generated_sql.CreateOutboxMessageParams{
		EventType: eventType,
		Payload:   payload,
	})
	return err
}
