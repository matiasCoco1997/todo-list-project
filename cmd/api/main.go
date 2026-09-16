package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matiasCoco1997/todo-list/internal/model"
)

func main() {
	http.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		items := []model.Item{
			{
				ID:             1,
				Nombre:         "Leche",
				Cantidad:       2,
				Categoria:      "Lácteos",
				Comprado:       false,
				PrecioEstimado: 1500,
			},
			{
				ID:             2,
				Nombre:         "Papel higiénico",
				Cantidad:       1,
				Categoria:      "Limpieza",
				Comprado:       false,
				PrecioEstimado: 5000,
			},
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(items)
		if err != nil {
			http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Servidor escuchando en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
