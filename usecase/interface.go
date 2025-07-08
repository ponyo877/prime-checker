package usecase

import (
	"context"

	"github.com/ponyo877/product-expiry-tracker/model"
)

type Repository interface {
	CreatePrimeCheck(ctx context.Context, userID int32, numberText string) (*model.PrimeCheck, error)
	GetPrimeCheck(ctx context.Context, id int32) (*model.PrimeCheck, error)
	ListPrimeChecks(ctx context.Context) ([]*model.PrimeCheck, error)
}
