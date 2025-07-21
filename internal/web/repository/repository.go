package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ponyo877/product-expiry-tracker/db/generated_sql"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
	"github.com/ponyo877/product-expiry-tracker/internal/web/model"
	"github.com/ponyo877/product-expiry-tracker/internal/web/usecase"
)

func convertNullStringToPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func convertNullBoolToPtr(nb sql.NullBool) *bool {
	if !nb.Valid {
		return nil
	}
	return &nb.Bool
}

type Repository struct {
	db      *sql.DB
	queries *generated_sql.Queries
}

func NewRepository(db *sql.DB) usecase.Repository {
	return &Repository{
		db:      db,
		queries: generated_sql.New(db),
	}
}

func (r *Repository) GetPrimeCheck(ctx context.Context, id int32) (*model.PrimeCheck, error) {
	test, err := r.queries.GetPrimeCheck(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewPrimeCheckWithExtras(
		test.ID, 
		test.UserID, 
		test.NumberText, 
		test.CreatedAt, 
		test.UpdatedAt,
		convertNullStringToPtr(test.TraceID),
		convertNullStringToPtr(test.MessageID),
		convertNullBoolToPtr(test.IsPrime),
		convertNullStringToPtr(test.Status),
	), nil
}

func (r *Repository) ListPrimeChecks(ctx context.Context) ([]*model.PrimeCheck, error) {
	tests, err := r.queries.ListPrimeChecks(ctx)
	if err != nil {
		return nil, err
	}

	result := []*model.PrimeCheck{}
	for _, test := range tests {
		result = append(result, model.NewPrimeCheckWithExtras(
			test.ID, 
			test.UserID, 
			test.NumberText, 
			test.CreatedAt, 
			test.UpdatedAt,
			convertNullStringToPtr(test.TraceID),
			convertNullStringToPtr(test.MessageID),
			convertNullBoolToPtr(test.IsPrime),
			convertNullStringToPtr(test.Status),
		))
	}
	return result, nil
}

func (r *Repository) CreatePrimeCheckWithMessage(ctx context.Context, userID int32, numberText string) (*model.PrimeCheck, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	txQueries := r.queries.WithTx(tx)

	result, err := txQueries.CreatePrimeCheck(ctx, generated_sql.CreatePrimeCheckParams{
		UserID:     userID,
		NumberText: numberText,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Create message for prime check worker with trace context
	payload := &message.PrimeCheckPayload{
		RequestID:  int32(id),
		UserID:     userID,
		NumberText: numberText,
	}

	msg, err := message.NewMessageWithTraceContext(ctx, message.MessageTypePrimeCheck, payload)
	if err != nil {
		return nil, err
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	// Save message to outbox
	if _, err := txQueries.CreateOutboxMessage(ctx, generated_sql.CreateOutboxMessageParams{
		EventType: string(message.MessageTypePrimeCheck),
		Payload:   msgBytes,
	}); err != nil {
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Return the created prime check
	check, err := r.queries.GetPrimeCheck(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return model.NewPrimeCheckWithExtras(
		check.ID, 
		check.UserID, 
		check.NumberText, 
		check.CreatedAt, 
		check.UpdatedAt,
		convertNullStringToPtr(check.TraceID),
		convertNullStringToPtr(check.MessageID),
		convertNullBoolToPtr(check.IsPrime),
		convertNullStringToPtr(check.Status),
	), nil
}
