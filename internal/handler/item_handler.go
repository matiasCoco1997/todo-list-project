package handler

import (
	"encoding/json"
	"net/http"

	"github.com/matiasCoco1997/todo-list/internal/service"
)

type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAllItems(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener los items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
