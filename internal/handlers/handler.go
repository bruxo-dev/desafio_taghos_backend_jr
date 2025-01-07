package handlers

import (
	"desafio_taghos/internal/models"
	"desafio_taghos/internal/services"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) (*Handler, error) {
	return &Handler{service: service}, nil
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.FetchItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.AddItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
