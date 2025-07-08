package usecase

import (
	"context"

	"github.com/ponyo877/product-expiry-tracker/model"
)

type Usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) CreatePrimeCheck(ctx context.Context, userID int32, numberText string) (*model.PrimeCheck, error) {
	return u.repo.CreatePrimeCheck(ctx, userID, numberText)
}

func (u *Usecase) GetPrimeCheck(ctx context.Context, id int32) (*model.PrimeCheck, error) {
	return u.repo.GetPrimeCheck(ctx, id)
}

func (u *Usecase) ListPrimeChecks(ctx context.Context) ([]*model.PrimeCheck, error) {
	return u.repo.ListPrimeChecks(ctx)
}
