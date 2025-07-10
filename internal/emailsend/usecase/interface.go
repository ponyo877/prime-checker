package usecase


type EmailSender interface {
	SendPrimeCheckResult(to, numberText string, isPrime bool) error
}