package repository

import (
	"context"
	"database/sql"

	"github.com/ponyo877/product-expiry-tracker/db/generated_sql"
	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/usecase"
)

type PrimeCheckRepository struct {
	db      *sql.DB
	queries *generated_sql.Queries
}

func NewPrimeCheckRepository(db *sql.DB) usecase.PrimeCheckRepository {
	return &PrimeCheckRepository{
		db:      db,
		queries: generated_sql.New(db),
	}
}

func (r *PrimeCheckRepository) UpdatePrimeCheckResult(ctx context.Context, requestID int32, traceID, messageID string, isPrime bool, status string) error {
	var traceIDPtr *string
	var messageIDPtr *string
	
	if traceID != "" {
		traceIDPtr = &traceID
	}
	if messageID != "" {
		messageIDPtr = &messageID
	}

	return r.queries.UpdatePrimeCheckResult(ctx, generated_sql.UpdatePrimeCheckResultParams{
		TraceID:   convertStringPtrToNullString(traceIDPtr),
		MessageID: convertStringPtrToNullString(messageIDPtr),
		IsPrime:   sql.NullBool{Bool: isPrime, Valid: true},
		Status:    sql.NullString{String: status, Valid: true},
		ID:        requestID,
	})
}

func convertStringPtrToNullString(ptr *string) sql.NullString {
	if ptr == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *ptr, Valid: true}
}