package adapter

import (
	"context"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"github.com/ponyo877/prime-checker/internal/web/usecase"
	"github.com/ponyo877/prime-checker/openapi"
)

type handler struct {
	usecase *usecase.Usecase
}

func NewHandler(uc *usecase.Usecase) *handler {
	return &handler{usecase: uc}
}

func (h *handler) PrimeChecksCreate(ctx context.Context, req *openapi.PrimeCheckInput) (r *openapi.PrimeCheck, _ error) {
	tracer := otel.Tracer("web-server")
	ctx, span := tracer.Start(ctx, "PrimeChecksCreate")
	defer span.End()

	traceID := span.SpanContext().TraceID().String()
	log.Printf("Processing request with Trace ID: %s", traceID)

	span.SetAttributes(
		attribute.String("number", req.Number),
		attribute.String("operation", "create_prime_check"),
	)

	// TODO: Replace with actual user ID retrieval logic
	userID := int32(1)

	test, err := h.usecase.CreatePrimeCheckWithMessage(ctx, userID, req.Number)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(attribute.Int("request_id", int(test.ID())))

	// Set trace ID for the created prime check
	test.SetTraceID(traceID)
	test.SetStatus("processing")

	return &openapi.PrimeCheck{
		ID:        test.ID(),
		Number:    test.NumberText(),
		CreatedAt: test.CreatedAt(),
		TraceID:   convertStringPtrToOptString(test.TraceID()),
		MessageID: convertStringPtrToOptString(test.MessageID()),
		IsPrime:   convertBoolPtrToOptBool(test.IsPrime()),
		Status:    convertStringPtrToOptString(test.Status()),
	}, nil
}

func (h *handler) PrimeChecksGet(ctx context.Context, params openapi.PrimeChecksGetParams) (r *openapi.PrimeCheck, _ error) {
	test, err := h.usecase.GetPrimeCheck(ctx, params.RequestID)
	if err != nil {
		return nil, err
	}

	return &openapi.PrimeCheck{
		ID:        test.ID(),
		Number:    test.NumberText(),
		CreatedAt: test.CreatedAt(),
		TraceID:   convertStringPtrToOptString(test.TraceID()),
		MessageID: convertStringPtrToOptString(test.MessageID()),
		IsPrime:   convertBoolPtrToOptBool(test.IsPrime()),
		Status:    convertStringPtrToOptString(test.Status()),
	}, nil
}

func (h *handler) PrimeChecksList(ctx context.Context) (r *openapi.PrimeCheckList, _ error) {
	tracer := otel.Tracer("web-server")
	ctx, span := tracer.Start(ctx, "PrimeChecksList")
	defer span.End()

	tests, err := h.usecase.ListPrimeChecks(ctx)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(attribute.Int("results_count", len(tests)))

	items := make([]openapi.PrimeCheck, len(tests))
	for i, test := range tests {
		items[i] = openapi.PrimeCheck{
			ID:        test.ID(),
			Number:    test.NumberText(),
			CreatedAt: test.CreatedAt(),
			TraceID:   convertStringPtrToOptString(test.TraceID()),
			MessageID: convertStringPtrToOptString(test.MessageID()),
			IsPrime:   convertBoolPtrToOptBool(test.IsPrime()),
			Status:    convertStringPtrToOptString(test.Status()),
		}
	}

	return &openapi.PrimeCheckList{
		Items: items,
	}, nil
}

func (h *handler) SettingsCreate(ctx context.Context, req *openapi.Setting) (r *openapi.Setting, _ error) {
	return nil, nil
}

func (h *handler) SettingsGet(ctx context.Context) (r *openapi.Setting, _ error) {
	return nil, nil
}

func (h *handler) NewError(ctx context.Context, err error) *openapi.ErrorStatusCode {
	return &openapi.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: openapi.Error{
			Code:    500,
			Message: err.Error(),
		},
	}
}

func convertStringPtrToOptString(ptr *string) openapi.OptString {
	if ptr == nil {
		return openapi.OptString{}
	}
	return openapi.NewOptString(*ptr)
}

func convertBoolPtrToOptBool(ptr *bool) openapi.OptBool {
	if ptr == nil {
		return openapi.OptBool{}
	}
	return openapi.NewOptBool(*ptr)
}
