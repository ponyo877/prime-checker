package service

import (
	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/email"
	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/usecase"
)

type SMTPSender struct {
	emailSender *email.Sender
}

func NewSMTPSender(emailSender *email.Sender) usecase.EmailSender {
	return &SMTPSender{
		emailSender: emailSender,
	}
}

func (s *SMTPSender) SendPrimeCheckResult(to, numberText string, isPrime bool) error {
	return s.emailSender.SendPrimeCheckResult(to, numberText, isPrime)
}
