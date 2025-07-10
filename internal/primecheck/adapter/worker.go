package adapter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/model"
	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/usecase"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
)

type PrimeCheckWorker struct {
	usecase *usecase.PrimeCheckUsecase
}

func NewPrimeCheckWorker(usecase *usecase.PrimeCheckUsecase) *PrimeCheckWorker {
	return &PrimeCheckWorker{
		usecase: usecase,
	}
}

func (w *PrimeCheckWorker) HandleMessage(ctx context.Context, msg *message.Message) error {
	log.Printf("Processing prime check message: %s", msg.ID)

	payload, err := msg.UnmarshalPrimeCheckPayload()
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	request := model.NewPrimeRequest(payload.RequestID, payload.UserID, payload.NumberText, time.Now())

	_, err = w.usecase.ProcessPrimeRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to process prime request: %w", err)
	}

	return nil
}
