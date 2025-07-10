package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ponyo877/product-expiry-tracker/internal/shared/config"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/infrastructure"
	"github.com/ponyo877/product-expiry-tracker/internal/web/adapter"
	"github.com/ponyo877/product-expiry-tracker/internal/web/repository"
	"github.com/ponyo877/product-expiry-tracker/internal/web/usecase"
	"github.com/ponyo877/product-expiry-tracker/openapi"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target ./openapi -package openapi --clean typespec/tsp-output/@typespec/openapi3/openapi.yaml

func main() {
	// Load configurations
	dbConfig := config.LoadDatabaseConfig()

	// Initialize infrastructure
	db, err := infrastructure.NewDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	uc := usecase.NewUseCase(repo)
	h := adapter.NewHandler(uc)
	srv, err := openapi.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}

	httpPort := ":8080"
	fmt.Printf("Starting web server on %s\n", httpPort)

	if err := http.ListenAndServe(httpPort, srv); err != nil {
		log.Fatal(err)
	}
}
