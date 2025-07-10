package adapter

import (
	"context"
	"log"
	"time"

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
			results, err := w.usecase.PublishPendingMessages(ctx)
			if err != nil {
				log.Printf("Error publishing pending messages: %v", err)
				continue
			}

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