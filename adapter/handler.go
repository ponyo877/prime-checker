package adapter

import (
	"context"
	"net/http"
	"time"

	"github.com/ponyo877/product-expiry-tracker/openapi"
	"github.com/ponyo877/product-expiry-tracker/usecase"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target ./openapi -package openapi --clean typespec/tsp-output/@typespec/openapi3/openapi.yaml

type handler struct {
	usecase *usecase.Usecase
}

func NewHandler(uc *usecase.Usecase) *handler {
	return &handler{usecase: uc}
}

// POST /primality-tests
func (h *handler) PrimeTestsCreate(ctx context.Context, req *openapi.PrimeTestRequest) (r *openapi.PrimeTest, _ error) {
	return nil, nil
}

// GET /primality-tests/{request_id}
func (h *handler) PrimeTestsGet(ctx context.Context, params openapi.PrimeTestsGetParams) (r *openapi.PrimeTest, _ error) {
	return nil, nil
}

// GET /primality-tests
func (h *handler) PrimeTestsList(ctx context.Context) (r *openapi.PrimeTestList, _ error) {
	return &openapi.PrimeTestList{
		Items: []openapi.PrimeTest{
			{
				ID:        1,
				Number:    "7",
				CreatedAt: time.Now(),
			},
		},
	}, nil
}

// POST /settings
func (h *handler) SettingsCreate(ctx context.Context, req *openapi.Setting) (r *openapi.Setting, _ error) {
	return nil, nil
}

// GET /settings
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
