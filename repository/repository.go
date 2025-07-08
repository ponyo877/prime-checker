package repository

import (
	"context"
	"database/sql"

	"github.com/ponyo877/product-expiry-tracker/db/generated_sql"
	"github.com/ponyo877/product-expiry-tracker/model"
	"github.com/ponyo877/product-expiry-tracker/usecase"
)

type Repository struct {
	db      *sql.DB
	queries *generated_sql.Queries
}

func NewRepository(db *sql.DB) usecase.Repository {
	return &Repository{
		db:      db,
		queries: generated_sql.New(db),
	}
}

func (r *Repository) CreatePrimeCheck(ctx context.Context, userID int32, numberText string) (*model.PrimeCheck, error) {
	result, err := r.queries.CreatePrimeCheck(ctx, generated_sql.CreatePrimeCheckParams{
		UserID:     userID,
		NumberText: numberText,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	test, err := r.queries.GetPrimeCheck(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return model.NewPrimeCheck(test.ID, test.UserID, test.NumberText, test.CreatedAt, test.UpdatedAt), nil
}

func (r *Repository) GetPrimeCheck(ctx context.Context, id int32) (*model.PrimeCheck, error) {
	test, err := r.queries.GetPrimeCheck(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewPrimeCheck(test.ID, test.UserID, test.NumberText, test.CreatedAt, test.UpdatedAt), nil
}

func (r *Repository) ListPrimeChecks(ctx context.Context) ([]*model.PrimeCheck, error) {
	tests, err := r.queries.ListPrimeChecks(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.PrimeCheck, len(tests))
	for i, test := range tests {
		result[i] = model.NewPrimeCheck(test.ID, test.UserID, test.NumberText, test.CreatedAt, test.UpdatedAt)
	}
	return result, nil
}
