package model

import "time"

type PrimeResult struct {
	requestID       int32
	userID          int32
	numberText      string
	isPrime         bool
	calculatedAt    time.Time
	calculationTime time.Duration
}

func NewPrimeResult(requestID, userID int32, numberText string, isPrime bool, now time.Time, calculationTime time.Duration) *PrimeResult {
	return &PrimeResult{
		requestID:       requestID,
		userID:          userID,
		numberText:      numberText,
		isPrime:         isPrime,
		calculatedAt:    now,
		calculationTime: calculationTime,
	}
}

func (p *PrimeResult) RequestID() int32 {
	return p.requestID
}

func (p *PrimeResult) UserID() int32 {
	return p.userID
}

func (p *PrimeResult) NumberText() string {
	return p.numberText
}

func (p *PrimeResult) IsPrime() bool {
	return p.isPrime
}

func (p *PrimeResult) CalculatedAt() time.Time {
	return p.calculatedAt
}

func (p *PrimeResult) CalculationTime() time.Duration {
	return p.calculationTime
}
