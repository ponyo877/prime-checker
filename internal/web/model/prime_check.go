package model

import (
	"time"
)

type PrimeCheck struct {
	id         int32
	userID     int32
	numberText string
	createdAt  time.Time
	updatedAt  time.Time
	traceID    *string
	messageID  *string
	isPrime    *bool
	status     *string
}

func NewPrimeCheck(id, userID int32, numberText string, createdAt, updatedAt time.Time) *PrimeCheck {
	return &PrimeCheck{
		id:         id,
		userID:     userID,
		numberText: numberText,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
		traceID:    nil,
		messageID:  nil,
		isPrime:    nil,
		status:     nil,
	}
}

func NewPrimeCheckWithExtras(id, userID int32, numberText string, createdAt, updatedAt time.Time, traceID, messageID *string, isPrime *bool, status *string) *PrimeCheck {
	return &PrimeCheck{
		id:         id,
		userID:     userID,
		numberText: numberText,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
		traceID:    traceID,
		messageID:  messageID,
		isPrime:    isPrime,
		status:     status,
	}
}

func (p *PrimeCheck) ID() int32 {
	return p.id
}

func (p *PrimeCheck) UserID() int32 {
	return p.userID
}

func (p *PrimeCheck) NumberText() string {
	return p.numberText
}

func (p *PrimeCheck) CreatedAt() time.Time {
	return p.createdAt
}

func (p *PrimeCheck) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *PrimeCheck) TraceID() *string {
	return p.traceID
}

func (p *PrimeCheck) MessageID() *string {
	return p.messageID
}

func (p *PrimeCheck) IsPrime() *bool {
	return p.isPrime
}

func (p *PrimeCheck) Status() *string {
	return p.status
}

func (p *PrimeCheck) SetTraceID(traceID string) {
	p.traceID = &traceID
}

func (p *PrimeCheck) SetMessageID(messageID string) {
	p.messageID = &messageID
}

func (p *PrimeCheck) SetIsPrime(isPrime bool) {
	p.isPrime = &isPrime
}

func (p *PrimeCheck) SetStatus(status string) {
	p.status = &status
}
