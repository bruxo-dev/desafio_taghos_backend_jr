package routes

import (
	"desafio_taghos/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.GET("/items", func(context *gin.Context) {
		handler.GetItems(context.Writer, context.Request)
	})
	router.POST("/items", func(context *gin.Context) {
		handler.CreateItem(context.Writer, context.Request)
	})
}
