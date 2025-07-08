package model

import (
	"time"
)

type Book struct {
	id        int32
	title     string
	author    string
	content   string
	createdAt time.Time
}

func NewBook(id int32, title, author, content string, createdAt time.Time) *Book {
	return &Book{
		id:        id,
		title:     title,
		author:    author,
		content:   content,
		createdAt: createdAt,
	}
}

func (b *Book) ID() int32 {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Content() string {
	return b.content
}

func (b *Book) CreatedAt() time.Time {
	return b.createdAt
}
