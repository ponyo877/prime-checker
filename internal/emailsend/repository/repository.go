package repository

import (
	"fmt"
	"net/smtp"

	"github.com/ponyo877/prime-checker/internal/emailsend/usecase"
)

type emailRepository struct {
	smtpHost string
	smtpPort string
	username string
}

func NewEmailRepository(smtpHost, smtpPort, username string) usecase.EmailRepository {
	return &emailRepository{
		smtpHost,
		smtpPort,
		username,
	}
}

func (r *emailRepository) SendEmail(to, subject, body, messageID string) error {
	addr := fmt.Sprintf("%s:%s", r.smtpHost, r.smtpPort)
	from := r.username
	
	var msg []byte
	if messageID != "" {
		msg = []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nMessage-ID: %s\r\n\r\n%s", to, subject, messageID, body))
	} else {
		msg = []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))
	}

	if err := smtp.SendMail(addr, nil, from, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
