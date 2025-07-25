package usecase

import (
	"context"

	"github.com/ponyo877/prime-checker/internal/web/model"
)

type Repository interface {
	GetPrimeCheck(ctx context.Context, id int32) (*model.PrimeCheck, error)
	ListPrimeChecks(ctx context.Context) ([]*model.PrimeCheck, error)
	CreatePrimeCheckWithMessage(ctx context.Context, userID int32, numberText string) (*model.PrimeCheck, error)
}
