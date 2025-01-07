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

func (s *Service) FetchUsers() ([]models.User, error) {
	rows, err := s.db.Query("SELECT id, username, password, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			return nil, err
		}
		user.Books, err = s.FetchBooksByUserID(user.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Service) AddUser(user models.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", user.Username, user.Password, user.Email)
	return err
}

func (s *Service) FetchBooks() ([]models.Book, error) {
	rows, err := s.db.Query("SELECT id, title, category, author, synopsis, user_id FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis, &book.UserID); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *Service) AddBook(book models.Book) error {
	_, err := s.db.Exec("INSERT INTO books (title, category, author, synopsis, user_id) VALUES ($1, $2, $3, $4, $5)", book.Title, book.Category, book.Author, book.Synopsis, book.UserID)
	return err
}

func (s *Service) FetchBooksByUserID(userID int) ([]models.Book, error) {
	rows, err := s.db.Query("SELECT id, title, category, author, synopsis, user_id FROM books WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis, &book.UserID); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
