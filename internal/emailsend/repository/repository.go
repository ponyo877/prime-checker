package repository

import (
	"fmt"
	"net/smtp"

	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/usecase"
)

type emailRepository struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewEmailRepository(smtpHost, smtpPort, username, password string) usecase.EmailRepository {
	return &emailRepository{
		smtpHost,
		smtpPort,
		username,
		password,
	}
}

func (r *emailRepository) SendEmail(to, subject, body string) error {
	if r.username == "" || r.password == "" {
		return fmt.Errorf("SMTP credentials not configured")
	}

	addr := fmt.Sprintf("%s:%s", r.smtpHost, r.smtpPort)
	auth := smtp.PlainAuth("", r.username, r.password, r.smtpHost)
	from := r.username
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

	if err := smtp.SendMail(addr, auth, from, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
