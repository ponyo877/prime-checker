package model

import "time"

type PublicationStatus string

const (
	PublicationStatusSuccess PublicationStatus = "success"
	PublicationStatusFailed  PublicationStatus = "failed"
	PublicationStatusSkipped PublicationStatus = "skipped"
)

type PublicationResult struct {
	messageID int32
	status    PublicationStatus
	error     error
	timestamp time.Time
}

func NewPublicationResult(messageID int32, status PublicationStatus, err error, now time.Time) *PublicationResult {
	return &PublicationResult{
		messageID: messageID,
		status:    status,
		error:     err,
		timestamp: now,
	}
}

func (p *PublicationResult) MessageID() int32 {
	return p.messageID
}

func (p *PublicationResult) Status() PublicationStatus {
	return p.status
}

func (p *PublicationResult) Error() error {
	return p.error
}

func (p *PublicationResult) Timestamp() time.Time {
	return p.timestamp
}

func (p *PublicationResult) IsSuccess() bool {
	return p.status == PublicationStatusSuccess
}
