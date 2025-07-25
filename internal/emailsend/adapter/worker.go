package adapter

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"

	"github.com/ponyo877/prime-checker/internal/emailsend/model"
	"github.com/ponyo877/prime-checker/internal/emailsend/usecase"
	"github.com/ponyo877/prime-checker/internal/shared/message"
)

type EmailSendWorker struct {
	usecase *usecase.EmailSendUsecase
}

func NewEmailSendWorker(usecase *usecase.EmailSendUsecase) *EmailSendWorker {
	return &EmailSendWorker{
		usecase: usecase,
	}
}

func (w *EmailSendWorker) HandleMessage(ctx context.Context, msg *message.Message) error {
	// Extract trace context from message
	ctx = msg.ExtractTraceContext(ctx)
	
	tracer := otel.Tracer("email-send-worker")
	ctx, span := tracer.Start(ctx, "HandleEmailSendMessage")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	log.Printf("Processing email send message: %s with Trace ID: %s", msg.ID, traceID)

	payload, err := msg.UnmarshalEmailSendPayload()
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	request := model.NewEmailRequest(
		payload.RequestID,
		payload.UserID,
		payload.Email,
		payload.Subject,
		payload.Body,
		payload.IsPrime,
		payload.NumberText,
		payload.MessageID,
	)

	result, err := w.usecase.SendPrimeCheckResult(request)
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	if result.IsSuccess() {
		log.Printf("Email sent successfully for request ID %d", result.RequestID())
	} else {
		log.Printf("Email sending failed for request ID %d: %v", result.RequestID(), result.Error())
	}

	return nil
}
