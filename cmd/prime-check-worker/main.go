package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ponyo877/prime-checker/internal/primecheck/adapter"
	"github.com/ponyo877/prime-checker/internal/primecheck/repository"
	"github.com/ponyo877/prime-checker/internal/primecheck/usecase"
	"github.com/ponyo877/prime-checker/internal/shared/config"
	"github.com/ponyo877/prime-checker/internal/shared/infrastructure"
)

func main() {
	// Initialize tracing
	tracingConfig := infrastructure.LoadTracingConfig("prime-check-worker")
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
	primeCheckRepo := repository.NewPrimeCheckRepository(db)
	calculator := repository.NewPrimeCalculator()
	publisher := repository.NewResultPublisher(outboxRepo)
	primeUsecase := usecase.NewPrimeCheckUsecase(calculator, publisher, primeCheckRepo)
	worker := adapter.NewPrimeCheckWorker(primeUsecase)

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

	log.Println("Starting prime check worker...")
	if err := natsBroker.Subscribe(ctx, "primecheck", worker.HandleMessage); err != nil && err != context.Canceled {
		log.Fatal("Prime check worker failed:", err)
	}

	log.Println("Prime check worker shutdown complete")
}
