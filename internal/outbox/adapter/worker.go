package adapter

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/ponyo877/product-expiry-tracker/internal/outbox/usecase"
)

type OutboxWorker struct {
	usecase *usecase.OutboxPublishingUsecase
}

func NewOutboxWorker(usecase *usecase.OutboxPublishingUsecase) *OutboxWorker {
	return &OutboxWorker{
		usecase: usecase,
	}
}

func (w *OutboxWorker) Start(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	log.Println("Starting outbox worker...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Outbox worker stopped")
			return ctx.Err()
		case <-ticker.C:
			tracer := otel.Tracer("outbox-publisher")
			tickerCtx, span := tracer.Start(ctx, "PublishPendingMessages")
			
			results, err := w.usecase.PublishPendingMessages(tickerCtx)
			if err != nil {
				span.RecordError(err)
				log.Printf("Error publishing pending messages: %v", err)
				span.End()
				continue
			}
			
			span.End()

			successCount := 0
			failedCount := 0
			for _, result := range results {
				if result.IsSuccess() {
					successCount++
				} else {
					failedCount++
				}
			}

			if len(results) > 0 {
				log.Printf("Published %d messages successfully, %d failed", successCount, failedCount)
			}
		}
	}
}
