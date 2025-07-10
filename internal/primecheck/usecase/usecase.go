package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/model"
)

type PrimeCheckUsecase struct {
	calculator PrimeCalculator
	publisher  ResultPublisher
}

func NewPrimeCheckUsecase(calculator PrimeCalculator, publisher ResultPublisher) *PrimeCheckUsecase {
	return &PrimeCheckUsecase{
		calculator: calculator,
		publisher:  publisher,
	}
}

func (u *PrimeCheckUsecase) ProcessPrimeRequest(ctx context.Context, request *model.PrimeRequest) (*model.PrimeResult, error) {
	log.Printf("Processing prime check request for number: %s", request.NumberText())

	startTime := time.Now()
	isPrime, err := u.calculator.Calculate(request.NumberText())
	calculationTime := time.Since(startTime)

	if err != nil {
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

	// Publish result for email notification
	if err := u.publisher.PublishEmailMessage(ctx, result); err != nil {
		log.Printf("Failed to publish email message: %v", err)
		// Don't return error here - the calculation was successful
	}

	return result, nil
}
