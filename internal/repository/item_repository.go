package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matiasCoco1997/todo-list/internal/model"
)

type ItemRepository struct {
	db *pgxpool.Pool
}

func NewItemRepository(db *pgxpool.Pool) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) GetAll(ctx context.Context) ([]model.Item, error) {
	query := `SELECT id, nombre, cantidad::float4, categoria, comprado, precio_estimado::float8 FROM items`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar query: %w", err)
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		err := rows.Scan(
			&item.ID,
			&item.Nombre,
			&item.Cantidad,
			&item.Categoria,
			&item.Comprado,
			&item.PrecioEstimado,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}
