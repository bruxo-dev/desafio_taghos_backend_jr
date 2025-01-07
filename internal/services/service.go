package services

import (
	"database/sql"
	"desafio_taghos/internal/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) FetchBooks() ([]models.Book, error) {
	rows, err := s.db.Query("SELECT id, title, category, author, synopsis FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *Service) FetchBookByID(id int) (*models.Book, error) {
	row := s.db.QueryRow("SELECT id, title, category, author, synopsis FROM books WHERE id = $1", id)
	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if book not found
		}
		return nil, err
	}
	return &book, nil
}

func (s *Service) AddBook(book models.Book) error {
	_, err := s.db.Exec("INSERT INTO books (title, category, author, synopsis) VALUES ($1, $2, $3, $4)", book.Title, book.Category, book.Author, book.Synopsis)
	return err
}

func (s *Service) UpdateBook(book models.Book) error {
	_, err := s.db.Exec("UPDATE books SET title = $1, category = $2, author = $3, synopsis = $4 WHERE id = $5", book.Title, book.Category, book.Author, book.Synopsis, book.ID)
	return err
}

func (s *Service) DeleteBook(id int) error {
	_, err := s.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
