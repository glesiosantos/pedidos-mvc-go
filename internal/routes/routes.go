package routes

import "github.com/gin-gonic/gin"

func Novo() *gin.Engine {
	routes := gin.Default()

	routes.GET("/heath", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	return routes
}
