package usecase

import "github.com/ponyo877/product-expiry-tracker/model"

type Repository interface {
	ListBooksByWord(word string) ([]*model.Book, error)
}
