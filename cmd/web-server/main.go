package main

import (
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/ponyo877/product-expiry-tracker/internal/shared/config"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/infrastructure"
	"github.com/ponyo877/product-expiry-tracker/internal/web/adapter"
	"github.com/ponyo877/product-expiry-tracker/internal/web/repository"
	"github.com/ponyo877/product-expiry-tracker/internal/web/usecase"
	"github.com/ponyo877/product-expiry-tracker/openapi"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target ../../openapi --package openapi --clean ../../typespec/tsp-output/@typespec/openapi3/openapi.yaml

func main() {
	// Initialize tracing
	tracingConfig := infrastructure.LoadTracingConfig("web-server")
	tp, err := infrastructure.InitTracing(tracingConfig)
	if err != nil {
		log.Fatal("Failed to initialize tracing:", err)
	}
	defer infrastructure.ShutdownTracing(tp)

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

	// CORS middleware
	corsHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	// Wrap server with CORS and OpenTelemetry instrumentation
	handler := corsHandler(otelhttp.NewHandler(srv, "web-server"))

	httpPort := ":8080"
	fmt.Printf("Starting web server on %s\n", httpPort)

	if err := http.ListenAndServe(httpPort, handler); err != nil {
		log.Fatal(err)
	}
}
