package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/adapter"
	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/email"
	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/service"
	"github.com/ponyo877/product-expiry-tracker/internal/emailsend/usecase"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/config"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/infrastructure"
)

func main() {
	// Load configurations
	msgConfig := config.LoadMessagingConfig()

	// Initialize infrastructure
	natsBroker, err := infrastructure.NewMessageBroker(msgConfig)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer natsBroker.Close()

	// Create dependencies (DI)
	emailSender := email.NewSender()
	smtpSender := service.NewSMTPSender(emailSender)
	emailUsecase := usecase.NewEmailSendUsecase(smtpSender)
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
