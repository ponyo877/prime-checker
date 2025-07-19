package infrastructure

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type TraceContext struct {
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
	Flags   string `json:"flags"`
}

type MessageWithTrace struct {
	TraceContext *TraceContext `json:"trace_context,omitempty"`
	Data         interface{}   `json:"data"`
}

func InjectTraceToMessage(ctx context.Context, data interface{}) ([]byte, error) {
	span := trace.SpanFromContext(ctx)
	
	var traceCtx *TraceContext
	if span.SpanContext().IsValid() {
		traceCtx = &TraceContext{
			TraceID: span.SpanContext().TraceID().String(),
			SpanID:  span.SpanContext().SpanID().String(),
			Flags:   span.SpanContext().TraceFlags().String(),
		}
	}

	msg := MessageWithTrace{
		TraceContext: traceCtx,
		Data:         data,
	}

	return json.Marshal(msg)
}

func ExtractTraceFromMessage(data []byte) (context.Context, interface{}, error) {
	var msg MessageWithTrace
	if err := json.Unmarshal(data, &msg); err != nil {
		return context.Background(), nil, err
	}

	ctx := context.Background()
	if msg.TraceContext != nil {
		// Create a new span context from the extracted trace data
		carrier := make(map[string]string)
		carrier["traceparent"] = formatTraceparent(msg.TraceContext)
		
		propagator := otel.GetTextMapPropagator()
		ctx = propagator.Extract(ctx, propagation.MapCarrier(carrier))
	}

	return ctx, msg.Data, nil
}

func formatTraceparent(tc *TraceContext) string {
	return "00-" + tc.TraceID + "-" + tc.SpanID + "-" + tc.Flags
}