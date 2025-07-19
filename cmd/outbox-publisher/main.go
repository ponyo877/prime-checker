package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ponyo877/product-expiry-tracker/internal/outbox/adapter"
	"github.com/ponyo877/product-expiry-tracker/internal/outbox/repository"
	"github.com/ponyo877/product-expiry-tracker/internal/outbox/usecase"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/config"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/infrastructure"
)

func main() {
	// Initialize tracing
	tracingConfig := infrastructure.LoadTracingConfig("outbox-publisher")
	tp, err := infrastructure.InitTracing(tracingConfig)
	if err != nil {
		log.Fatal("Failed to initialize tracing:", err)
	}
	defer infrastructure.ShutdownTracing(tp)

	// Load configurations
	dbConfig := config.LoadDatabaseConfig()
	msgConfig := config.LoadMessagingConfig()

	// Initialize infrastructure
	db, err := infrastructure.NewDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	natsBroker, err := infrastructure.NewMessageBroker(msgConfig)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer natsBroker.Close()

	// Create dependencies (DI)
	queries := infrastructure.NewQueries(db)
	outboxRepo := repository.NewOutboxRepository(queries)
	messagePublisher := adapter.NewMessagePublisher(natsBroker)
	outboxUsecase := usecase.NewOutboxPublishingUsecase(outboxRepo, messagePublisher)
	worker := adapter.NewOutboxWorker(outboxUsecase)

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

	log.Println("Starting outbox publisher...")
	if err := worker.Start(ctx); err != nil && err != context.Canceled {
		log.Fatal("Outbox publisher failed:", err)
	}

	log.Println("Outbox publisher shutdown complete")
}
