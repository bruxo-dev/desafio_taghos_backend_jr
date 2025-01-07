package main

import (
	"desafio_taghos/internal/handlers"
	"desafio_taghos/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	handler, err := handlers.NewHandler()
	if err != nil {
		log.Fatalf("Could not initialize handler: %s\n", err)
	}
	routes.SetupRoutes(engine, handler)
	log.Println("Starting server on :8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
