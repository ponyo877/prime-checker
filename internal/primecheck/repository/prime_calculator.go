package repository

import (
	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/model"
	"github.com/ponyo877/product-expiry-tracker/internal/primecheck/usecase"
)

type PrimeCalculator struct{}

func NewPrimeCalculator() usecase.PrimeCalculator {
	return &PrimeCalculator{}
}

func (c *PrimeCalculator) Calculate(numberText string) (bool, error) {
	checker, err := model.NewPrimeChecker(numberText)
	if err != nil {
		return false, err
	}

	return checker.IsPrime(), nil
}
