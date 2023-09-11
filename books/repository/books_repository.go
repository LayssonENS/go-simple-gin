package booksRepository

import (
	"database/sql"
	"time"

	"github.com/LayssonENS/go-simple-gin/domain"
)

const dateLayout = "2006-01-02"

type postgresBooksRepo struct {
	DB *sql.DB
}

// NewPostgresBooksRepository will create an implementation of Books.Repository
func NewPostgresBooksRepository(db *sql.DB) domain.BooksRepository {
	return &postgresBooksRepo{
		DB: db,
	}
}

// GetByID : Retrieves a books by ID from the Postgres repository
func (p *postgresBooksRepo) GetByID(id int64) (domain.Books, error) {
	var books domain.Books
	err := p.DB.QueryRow(
		"SELECT id, name, email, birth_date, created_at FROM books WHERE id = $1", id).Scan(
		&books.ID, &books.Name, &books.Email, &books.BirthDate, &books.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return books, domain.ErrRegistrationNotFound
		}
		return books, err
	}
	return books, nil
}

// CreateBooks : Inserts a new books into the Postgres repository using the provided books request data
func (p *postgresBooksRepo) CreateBooks(books domain.BooksRequest) error {
	date, _ := time.Parse(dateLayout, books.BirthDate)
	birthDate := date

	query := `INSERT INTO books (name, email, birth_date) VALUES ($1, $2, $3)`
	_, err := p.DB.Exec(query, books.Name, books.Email, birthDate)
	if err != nil {
		return err
	}

	return nil
}

// GetAllBooks : Retrieves all books data from the Postgres repository
func (p *postgresBooksRepo) GetAllBooks() ([]domain.Books, error) {
	var bookss []domain.Books

	rows, err := p.DB.Query("SELECT id, name, email, birth_date, created_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var books domain.Books
		err := rows.Scan(
			&books.ID,
			&books.Name,
			&books.Email,
			&books.BirthDate,
			&books.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		bookss = append(bookss, books)
	}

	if len(bookss) == 0 {
		return nil, domain.ErrRegistrationNotFound
	}

	return bookss, nil
}
