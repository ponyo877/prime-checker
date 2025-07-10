package adapter

import (
	"context"
	"net/http"

	"github.com/ponyo877/product-expiry-tracker/internal/web/usecase"
	"github.com/ponyo877/product-expiry-tracker/openapi"
)

type handler struct {
	usecase *usecase.Usecase
}

func NewHandler(uc *usecase.Usecase) *handler {
	return &handler{usecase: uc}
}

func (h *handler) PrimeChecksCreate(ctx context.Context, req *openapi.PrimeCheckInput) (r *openapi.PrimeCheck, _ error) {
	// TODO: Replace with actual user ID retrieval logic
	userID := int32(1)

	test, err := h.usecase.CreatePrimeCheckWithMessage(ctx, userID, req.Number)
	if err != nil {
		return nil, err
	}

	return &openapi.PrimeCheck{
		ID:        test.ID(),
		Number:    test.NumberText(),
		CreatedAt: test.CreatedAt(),
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
	}, nil
}

func (h *handler) PrimeChecksList(ctx context.Context) (r *openapi.PrimeCheckList, _ error) {
	tests, err := h.usecase.ListPrimeChecks(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]openapi.PrimeCheck, len(tests))
	for i, test := range tests {
		items[i] = openapi.PrimeCheck{
			ID:        test.ID(),
			Number:    test.NumberText(),
			CreatedAt: test.CreatedAt(),
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
