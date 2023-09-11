package usecase

import (
	"github.com/LayssonENS/go-simple-gin/domain"
)

type booksUseCase struct {
	booksRepository domain.BooksRepository
}

func NewBooksUseCase(booksRepository domain.BooksRepository) domain.BooksUseCase {
	return &booksUseCase{
		booksRepository: booksRepository,
	}
}

func (a *booksUseCase) GetByID(id int64) (domain.Books, error) {
	books, err := a.booksRepository.GetByID(id)
	if err != nil {
		return books, err
	}

	return books, nil
}

func (a *booksUseCase) CreateBooks(books domain.Books) error {
	err := a.booksRepository.CreateBooks(books)
	if err != nil {
		return err
	}

	return nil
}

func (a *booksUseCase) GetAllBooks() ([]domain.Books, error) {
	books, err := a.booksRepository.GetAllBooks()
	if err != nil {
		return books, err
	}

	return books, nil
}
