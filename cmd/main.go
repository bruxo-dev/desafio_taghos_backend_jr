package main

import (
	"database/sql"
	"desafio_taghos/internal/handlers"
	"desafio_taghos/internal/routes"
	"desafio_taghos/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err)
	}
	defer db.Close()

	service := services.NewService(db)
	handler, err := handlers.NewHandler(service)
	if err != nil {
		log.Fatalf("Could not initialize handler: %s\n", err)
	}

	engine := gin.Default()
	routes.SetupRoutes(engine, handler)
	log.Println("Starting server on :8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
