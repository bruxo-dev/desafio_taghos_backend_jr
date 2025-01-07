package routes

import (
	"desafio_taghos/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.GET("/users", func(context *gin.Context) {
		handler.GetUsers(context.Writer, context.Request)
	})
	router.POST("/users", func(context *gin.Context) {
		handler.CreateUser(context.Writer, context.Request)
	})
	router.GET("/books", func(context *gin.Context) {
		handler.GetBooks(context.Writer, context.Request)
	})
	router.POST("/books", func(context *gin.Context) {
		handler.CreateBook(context.Writer, context.Request)
	})
}
