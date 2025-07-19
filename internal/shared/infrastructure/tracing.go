package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type TracingConfig struct {
	ServiceName string
	Environment string
	Host        string
	Port        string
}

func InitTracing(config TracingConfig) (*trace.TracerProvider, error) {
	ctx := context.Background()

	// Create OTLP HTTP exporter
	endpoint := fmt.Sprintf("%s:%s", config.Host, config.Port)

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithURLPath("/v1/traces"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	// Create resource
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Create trace provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()),
	)

	// Set global tracer provider
	otel.SetTracerProvider(tp)

	log.Printf("Tracing initialized for service: %s", config.ServiceName)
	return tp, nil
}

func LoadTracingConfig(serviceName string) TracingConfig {
	return TracingConfig{
		ServiceName: serviceName,
		Environment: "development",
		Host:        os.Getenv("JAEGER_HOST"),
		Port:        os.Getenv("JAEGER_PORT"),
	}
}

func ShutdownTracing(tp *trace.TracerProvider) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := tp.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down tracer provider: %v", err)
	}
}
