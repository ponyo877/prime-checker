package usecase

import (
	"fmt"
	"log"

	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/model"
)

type EmailSendUsecase struct {
	sender EmailSender
}

func NewEmailSendUsecase(sender EmailSender) *EmailSendUsecase {
	return &EmailSendUsecase{
		sender: sender,
	}
}

func (u *EmailSendUsecase) SendPrimeCheckResult(request *model.EmailRequest) (*model.SendResult, error) {
	log.Printf("Sending email to %s for request ID %d", request.Email(), request.RequestID())

	err := u.sender.SendPrimeCheckResult(request.Email(), request.NumberText(), request.IsPrime())
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return model.NewSendResult(request.RequestID(), model.SendStatusFailed, err), fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully for request ID %d", request.RequestID())
	return model.NewSendResult(request.RequestID(), model.SendStatusSuccess, nil), nil
}
