package model

import "time"

type PrimeRequest struct {
	requestID  int32
	userID     int32
	numberText string
	timestamp  time.Time
}

func NewPrimeRequest(requestID, userID int32, numberText string, now time.Time) *PrimeRequest {
	return &PrimeRequest{
		requestID:  requestID,
		userID:     userID,
		numberText: numberText,
		timestamp:  now,
	}
}

func (p *PrimeRequest) RequestID() int32 {
	return p.requestID
}

func (p *PrimeRequest) UserID() int32 {
	return p.userID
}

func (p *PrimeRequest) NumberText() string {
	return p.numberText
}

func (p *PrimeRequest) Timestamp() time.Time {
	return p.timestamp
}
