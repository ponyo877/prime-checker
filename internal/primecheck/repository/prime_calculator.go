package repository

import (
	"github.com/ponyo877/prime-checker/internal/primecheck/model"
	"github.com/ponyo877/prime-checker/internal/primecheck/usecase"
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
