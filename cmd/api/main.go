package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matiasCoco1997/todo-list/internal/handler"
	"github.com/matiasCoco1997/todo-list/internal/repository"
	"github.com/matiasCoco1997/todo-list/internal/service"
)

func main() {
	// 1. Cadena de conexión a PostgreSQL (coincide con docker-compose.yml)
	dbURL := "postgres://postgres:secretpassword@localhost:5432/todolist?sslmode=disable"

	// 2. Conectar a la base de datos
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v\n", err)
	}
	defer dbPool.Close()

	// Probar conexión
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v\n", err)
	}
	fmt.Println("Conexión exitosa a PostgreSQL")

	// 3. Inyección de dependencias (Database -> Repository -> Service -> Handler)
	repo := repository.NewItemRepository(dbPool)
	svc := service.NewItemService(repo)
	h := handler.NewItemHandler(svc)

	// 4. Configurar rutas
	http.HandleFunc("/api/items", h.GetItems)

	// 5. Iniciar servidor
	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v\n", err)
	}
}
