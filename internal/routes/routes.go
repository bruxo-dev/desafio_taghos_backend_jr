package routes

import (
	"desafio_taghos/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.GET("/books", handler.GetBooks)
	router.GET("/books/:id", handler.GetBookById)
	router.POST("/books", handler.CreateBook)
	router.PUT("/books/:id", handler.UpdateBookById)
	router.DELETE("/books/:id", handler.DeleteBookById)
}
