package usecase

import (
	"fmt"
	"log"

	"github.com/ponyo877/prime-checker/internal/emailsend/model"
)

type EmailSendUsecase struct {
	repo EmailRepository
}

func NewEmailSendUsecase(repo EmailRepository) *EmailSendUsecase {
	return &EmailSendUsecase{
		repo: repo,
	}
}

func (u *EmailSendUsecase) SendPrimeCheckResult(request *model.EmailRequest) (*model.SendResult, error) {
	log.Printf("Sending email to %s for request ID %d", request.Email(), request.RequestID())

	var subject, body string
	if request.IsPrime() {
		subject = fmt.Sprintf("Prime Check Result: %s is Prime!", request.NumberText())
		body = fmt.Sprintf("Good news! The number %s is a prime number.", request.NumberText())
	} else {
		subject = fmt.Sprintf("Prime Check Result: %s is not Prime", request.NumberText())
		body = fmt.Sprintf("The number %s is not a prime number.", request.NumberText())
	}

	if err := u.repo.SendEmail(request.Email(), subject, body, request.MessageID()); err != nil {
		log.Printf("Failed to send email: %v", err)
		return model.NewSendResult(request.RequestID(), model.SendStatusFailed, err), fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully for request ID %d", request.RequestID())
	return model.NewSendResult(request.RequestID(), model.SendStatusSuccess, nil), nil
}
