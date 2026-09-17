package service

import (
	"context"

	"github.com/matiasCoco1997/todo-list/internal/model"
	"github.com/matiasCoco1997/todo-list/internal/repository"
)

type ItemService struct {
	repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) GetAllItems(ctx context.Context) ([]model.Item, error) {
	// Acá más adelante pondremos validaciones o lógica de negocio
	return s.repo.GetAll(ctx)
}
