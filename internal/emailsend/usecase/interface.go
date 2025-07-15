package usecase

type EmailRepository interface {
	SendEmail(to, subject, body string) error
}
