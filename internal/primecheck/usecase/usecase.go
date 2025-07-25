package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/ponyo877/prime-checker/internal/primecheck/model"
)

type PrimeCheckUsecase struct {
	calculator PrimeCalculator
	publisher  ResultPublisher
	repository PrimeCheckRepository
}

func NewPrimeCheckUsecase(calculator PrimeCalculator, publisher ResultPublisher, repository PrimeCheckRepository) *PrimeCheckUsecase {
	return &PrimeCheckUsecase{
		calculator: calculator,
		publisher:  publisher,
		repository: repository,
	}
}

func (u *PrimeCheckUsecase) ProcessPrimeRequest(ctx context.Context, request *model.PrimeRequest) (*model.PrimeResult, error) {
	log.Printf("Processing prime check request for number: %s", request.NumberText())

	startTime := time.Now()
	isPrime, err := u.calculator.Calculate(request.NumberText())
	calculationTime := time.Since(startTime)

	if err != nil {
		// Update DB with failed status
		if updateErr := u.repository.UpdatePrimeCheckResult(ctx, request.RequestID(), "", "", false, "failed"); updateErr != nil {
			log.Printf("Failed to update prime check result in DB: %v", updateErr)
		}
		return nil, fmt.Errorf("failed to calculate prime: %w", err)
	}

	result := model.NewPrimeResult(
		request.RequestID(),
		request.UserID(),
		request.NumberText(),
		isPrime,
		time.Now(),
		calculationTime,
	)

	log.Printf("Prime check result for %s: %v (took %v)", request.NumberText(), isPrime, calculationTime)

	// Update DB with completed result
	traceID := getTraceIDFromContext(ctx)
	messageID := fmt.Sprintf("msg_%d_%d", request.RequestID(), time.Now().Unix())
	if err := u.repository.UpdatePrimeCheckResult(ctx, request.RequestID(), traceID, messageID, isPrime, "completed"); err != nil {
		log.Printf("Failed to update prime check result in DB: %v", err)
		// Continue with email publishing even if DB update fails
	}

	// Publish result for email notification
	if err := u.publisher.PublishEmailMessage(ctx, result, messageID); err != nil {
		log.Printf("Failed to publish email message: %v", err)
		// Don't return error here - the calculation was successful
	}

	return result, nil
}

func getTraceIDFromContext(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return span.SpanContext().TraceID().String()
	}
	return ""
}
