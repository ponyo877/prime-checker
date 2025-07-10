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
}

func NewPrimeCheck(id, userID int32, numberText string, createdAt, updatedAt time.Time) *PrimeCheck {
	return &PrimeCheck{
		id:         id,
		userID:     userID,
		numberText: numberText,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
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
