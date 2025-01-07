package handlers

import (
	"desafio_taghos/internal/models"
	"encoding/json"
	"net/http"
)

type Handler struct{}

func NewHandler() (*Handler, error) {
	// Initialize any dependencies here if needed
	return &Handler{}, nil
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch items
	items := []models.Item{} // Replace with actual fetching logic
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to create item
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
