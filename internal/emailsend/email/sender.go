package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Sender struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewSender() *Sender {
	return &Sender{
		smtpHost: getEnvOrDefault("SMTP_HOST", "localhost"),
		smtpPort: getEnvOrDefault("SMTP_PORT", "1025"),
		username: getEnvOrDefault("SMTP_USERNAME", "test@example.com"),
		password: getEnvOrDefault("SMTP_PASSWORD", "password"),
	}
}

func (s *Sender) SendEmail(to, subject, body string) error {
	if s.username == "" || s.password == "" {
		return fmt.Errorf("SMTP credentials not configured")
	}

	from := s.username
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	if err := smtp.SendMail(addr, auth, from, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}

func (s *Sender) SendPrimeCheckResult(to, numberText string, isPrime bool) error {
	var subject, body string
	if isPrime {
		subject = fmt.Sprintf("Prime Check Result: %s is Prime!", numberText)
		body = fmt.Sprintf("Good news! The number %s is a prime number.", numberText)
	} else {
		subject = fmt.Sprintf("Prime Check Result: %s is not Prime", numberText)
		body = fmt.Sprintf("The number %s is not a prime number.", numberText)
	}

	return s.SendEmail(to, subject, body)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
