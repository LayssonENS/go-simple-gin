package domain

import "time"

type Books struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  bool      `json:"isActive"`
}

type BooksUseCase interface {
	GetByID(id int64) (Books, error)
	CreateBooks(Books Books) error
	GetAllBooks() ([]Books, error)
}

type BooksRepository interface {
	GetByID(id int64) (Books, error)
	CreateBooks(Books Books) error
	GetAllBooks() ([]Books, error)
}
