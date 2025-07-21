package usecase

type EmailRepository interface {
	SendEmail(to, subject, body, messageID string) error
}
