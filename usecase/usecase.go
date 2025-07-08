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

func (u *Usecase) SearchBooks(ctx context.Context, word string) ([]*model.Book, error) {
	return u.repo.ListBooksByWord(word)
}
