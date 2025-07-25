package adapter

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/ponyo877/prime-checker/internal/primecheck/model"
	"github.com/ponyo877/prime-checker/internal/primecheck/usecase"
	"github.com/ponyo877/prime-checker/internal/shared/message"
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
	// Extract trace context from message
	ctx = msg.ExtractTraceContext(ctx)
	
	tracer := otel.Tracer("prime-check-worker")
	ctx, span := tracer.Start(ctx, "HandlePrimeCheckMessage")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	log.Printf("Processing prime check message: %s with Trace ID: %s", msg.ID, traceID)

	payload, err := msg.UnmarshalPrimeCheckPayload()
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	request := model.NewPrimeRequest(payload.RequestID, payload.UserID, payload.NumberText, time.Now())

	_, err = w.usecase.ProcessPrimeRequest(ctx, request)
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("failed to process prime request: %w", err)
	}

	return nil
}
