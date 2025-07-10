package model

import (
	"errors"
	"math/big"
)

const (
	// Number of iterations for primality test, same as UUID v4 collision probability
	iterations = 61
)

type PrimeChecker struct {
	number *big.Int
}

func NewPrimeChecker(numberText string) (*PrimeChecker, error) {
	bigNum := new(big.Int)
	if _, ok := bigNum.SetString(numberText, 10); !ok {
		return nil, errors.New("Invalid number format")
	}
	return &PrimeChecker{number: bigNum}, nil
}

func (c *PrimeChecker) IsPrime() bool {
	return c.number.ProbablyPrime(iterations)
}
