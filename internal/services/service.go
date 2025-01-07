package services

import (
	"database/sql"
	"desafio_taghos_backend_jr/internal/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) FetchItems() ([]models.Item, error) {
	rows, err := s.db.Query("SELECT id, name, description FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *Service) AddItem(item models.Item) error {
	_, err := s.db.Exec("INSERT INTO items (id, name, description) VALUES ($1, $2, $3)", item.ID, item.Name, item.Description)
	return err
}
