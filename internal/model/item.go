package model

type Item struct {
	ID             int     `json:"id"`
	Nombre         string  `json:"nombre"`
	Cantidad       float32 `json:"cantidad"`
	Categoria      string  `json:"categoria"`
	Comprado       bool    `json:"comprado"`
	PrecioEstimado float64 `json:"precioEstimado"`
}
