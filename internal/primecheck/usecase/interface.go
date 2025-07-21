package usecase

import (
	"context"
	"encoding/json"

	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/model"
)

type PrimeCalculator interface {
	Calculate(numberText string) (bool, error)
}

type ResultPublisher interface {
	PublishEmailMessage(ctx context.Context, result *model.PrimeResult) error
}

type OutboxRepository interface {
	CreateOutboxMessage(ctx context.Context, eventType string, payload json.RawMessage) error
}

type PrimeCheckRepository interface {
	UpdatePrimeCheckResult(ctx context.Context, requestID int32, traceID, messageID string, isPrime bool, status string) error
}
