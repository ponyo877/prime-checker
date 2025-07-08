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

func (r *Repository) ListBooksByWord(word string) ([]*model.Book, error) {
	ctx := context.Background()
	var books []*model.Book

	// Use LIKE search for short words (less than 2 characters due to ngram_token_size=2)
	if len([]rune(word)) < 2 {
		likePattern := word + "%"
		sqlcBooks, err := r.queries.SearchBooksByPattern(ctx, generated_sql.SearchBooksByPatternParams{
			Title:  likePattern,
			Author: likePattern,
		})
		if err != nil {
			return nil, err
		}

		for _, book := range sqlcBooks {
			books = append(books, model.NewBook(book.ID, book.Title, book.Author, "", book.CreatedAt))
		}
	} else {
		// sqlc not support full-text search so use raw SQL query
		query := "SELECT * FROM books WHERE MATCH (title, author) AGAINST (? IN BOOLEAN MODE)"
		rows, err := r.db.Query(query, word)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var book generated_sql.Book
			if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedAt, &book.CreatedAt); err != nil {
				return nil, err
			}
			books = append(books, model.NewBook(book.ID, book.Title, book.Author, "", book.CreatedAt))
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return books, nil
}
