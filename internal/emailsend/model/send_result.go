package model

import "time"

type SendStatus string

const (
	SendStatusSuccess SendStatus = "success"
	SendStatusFailed  SendStatus = "failed"
	SendStatusSkipped SendStatus = "skipped"
)

type SendResult struct {
	requestID int32
	status    SendStatus
	error     error
	sentAt    time.Time
}

func NewSendResult(requestID int32, status SendStatus, err error) *SendResult {
	return &SendResult{
		requestID: requestID,
		status:    status,
		error:     err,
		sentAt:    time.Now(),
	}
}

func (s *SendResult) RequestID() int32 {
	return s.requestID
}

func (s *SendResult) Status() SendStatus {
	return s.status
}

func (s *SendResult) Error() error {
	return s.error
}

func (s *SendResult) SentAt() time.Time {
	return s.sentAt
}

func (s *SendResult) IsSuccess() bool {
	return s.status == SendStatusSuccess
}
