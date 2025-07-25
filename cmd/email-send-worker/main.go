package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ponyo877/prime-checker/internal/emailsend/adapter"
	"github.com/ponyo877/prime-checker/internal/emailsend/repository"
	"github.com/ponyo877/prime-checker/internal/emailsend/usecase"
	"github.com/ponyo877/prime-checker/internal/shared/config"
	"github.com/ponyo877/prime-checker/internal/shared/infrastructure"
)

func main() {
	// Initialize tracing
	tracingConfig := infrastructure.LoadTracingConfig("email-send-worker")
	tp, err := infrastructure.InitTracing(tracingConfig)
	if err != nil {
		log.Fatal("Failed to initialize tracing:", err)
	}
	defer infrastructure.ShutdownTracing(tp)

	// Load configurations
	msgConfig := config.LoadMessagingConfig()

	// Initialize infrastructure
	natsBroker, err := infrastructure.NewMessageBroker(msgConfig)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer natsBroker.Close()

	// Create dependencies (DI)
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")

	emailRepo := repository.NewEmailRepository(smtpHost, smtpPort, username)
	emailUsecase := usecase.NewEmailSendUsecase(emailRepo)
	worker := adapter.NewEmailSendWorker(emailUsecase)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received shutdown signal")
		cancel()
	}()

	log.Println("Starting email send worker...")
	if err := natsBroker.Subscribe(ctx, "emailsend", worker.HandleMessage); err != nil && err != context.Canceled {
		log.Fatal("Email send worker failed:", err)
	}

	log.Println("Email send worker shutdown complete")
}
